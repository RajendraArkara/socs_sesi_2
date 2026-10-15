package deposit

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const depositsDir = "filebase/deposits"

const interestRate = 0.01

func Add(name string, amount float64, now time.Time) error {
	if err := os.MkdirAll(depositsDir, 0755); err != nil {
		return err
	}

	path := filepath.Join(depositsDir, name+".txt")

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	line := fmt.Sprintf("%.2f|%s\n", amount, now.Format(time.RFC3339))
	_, err = f.WriteString(line)
	return err
}

func AccrueAll(now time.Time) (int, error) {
	entries, err := os.ReadDir(depositsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}

	total := 0
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".txt") {
			continue
		}

		n, err := accrueFile(filepath.Join(depositsDir, e.Name()), now)
		if err != nil {
			return total, fmt.Errorf("%s: %w", e.Name(), err)
		}
		total += n
	}
	return total, nil
}

func accrueFile(path string, now time.Time) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	var out []string
	updated := 0

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			return 0, errors.New("format deposit tidak valid: " + line)
		}

		amount, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return 0, errors.New("jumlah deposit tidak valid: " + line)
		}

		last, err := time.Parse(time.RFC3339, parts[1])
		if err != nil {
			return 0, errors.New("timestamp deposit tidak valid: " + line)
		}

		minutes := int(now.Sub(last).Minutes())
		if minutes > 0 {
			amount = amount * math.Pow(1+interestRate, float64(minutes))
			last = last.Add(time.Duration(minutes) * time.Minute)
			updated++
		}

		out = append(out, fmt.Sprintf("%.2f|%s", amount, last.Format(time.RFC3339)))
	}

	content := strings.Join(out, "\n")
	if content != "" {
		content += "\n"
	}
	return updated, os.WriteFile(path, []byte(content), 0644)
}
