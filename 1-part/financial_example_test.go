package part_one_financial_example_test

import (
	part_one_financial_example "tdd/1-part"
	"testing"
)

func TestMultiply(t *testing.T) {
    tests := []struct {
        name     string
        a        int
        b        int
        expected int
    }{
        {
            name:     "positive numbers",
            a:        2,
            b:        3,
            expected: 6,
        },
        {
            name:     "multiply by zero",
            a:        5,
            b:        0,
            expected: 0,
        },
        {
            name:     "negative and positive",
            a:        -2,
            b:        3,
            expected: -6,
        },
        {
            name:     "large numbers",
            a:        10,
            b:        10,
            expected: 100,
        },
        {
            name:     "multiply by one",
            a:        7,
            b:        1,
            expected: 7,
        },
        {
            name:     "both negative",
            a:        -4,
            b:        -5,
            expected: 20,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := part_one_financial_example.UseMultiply(tt.a, tt.b)
            
            if result != tt.expected {
                t.Errorf("UseMultiply(%d, %d) = %d; want %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

func TestMultiplyFinancialScenarios(t *testing.T) {
    tests := []struct {
        name        string
        amount      int
        multiplier  int
        expected    int
        description string
    }{
        {
            name:        "monthly_to_yearly",
            amount:      100,
            multiplier:  12,
            expected:    1200,
            description: "monthly payment converted to yearly",
        },
        {
            name:        "installment_calculation",
            amount:      50,
            multiplier:  24,
            expected:    1200,
            description: "24 installments of 50",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := part_one_financial_example.UseMultiply(tt.amount, tt.multiplier)
            
            if result != tt.expected {
                t.Errorf("%s: UseMultiply(%d, %d) = %d; want %d", 
                    tt.description, tt.amount, tt.multiplier, result, tt.expected)
            }
        })
    }
}