package calculator_agent

import (
	"DistributedComputingServer/calculator_agent/calculator_multi"
	"database/sql"
	"fmt"
	shuntingYard "github.com/mgenware/go-shunting-yard"
	"strconv"
	"strings"
	"sync"
)

func AgentDo(input string, id int) error {
	infixTokens, err := shuntingYard.Scan(input)
	if err != nil {
		return err
	}
	postfixTokens, err := shuntingYard.Parse(infixTokens)
	if err != nil {
		return err
	}
	s := ""
	for _, t := range postfixTokens {
		if str, ok := t.Value.(string); ok {
			s += str + " "
		} else {
			s += strconv.Itoa(t.Value.(int)) + " "
		}
	}
	tokens := strings.Fields(s)
	stack := make([]int, 0)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return fmt.Errorf("not enough operands for addition")
			}
			var wg sync.WaitGroup
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			wg.Add(1)
			calculator_multi.Calculate(a, b, token, &stack, &wg)
			wg.Wait()

		default:
			if token == " " {
				continue
			}
			num, err := strconv.Atoi(token)
			if err != nil {
				return fmt.Errorf("invalid token %s", token)
			}

			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return fmt.Errorf("invalid expression: stack contains more than one result")
	}

	result := stack[0]

	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE tasks SET status = ?, expression = ? WHERE id = ?", "ready", result, id)
	if err != nil {
		return fmt.Errorf("failed to update task in database: %v", err)
	}

	return nil
}
