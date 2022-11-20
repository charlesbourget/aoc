package preparer

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://adventofcode.com/%d/day/%d/input"

func fetchInput(day int, year int, session string) ([]byte, error) {
	urlFormatted := fmt.Sprintf(url, year, day)

	client := http.Client{}
	req, err := http.NewRequest("GET", urlFormatted, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("cookie", fmt.Sprintf("session=%s", session))
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	var input []byte
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, err
		}

		input = bodyBytes
	} else {
		return []byte{}, fmt.Errorf("error while fetching input. Status code: %d", resp.StatusCode)
	}

	return input, nil
}
