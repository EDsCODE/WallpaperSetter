package scraper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// Item describes a Reddit item.
type Item struct {
	Title string
	URL   string
}

type response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func (i Item) String() string {

	return fmt.Sprintf("%s\n%s\n", i.Title, i.URL)
}

const UserAgent = "Reddit Wallpapesr (by /u/FutureBig)"

// Get fetches the most recent items in the subreddit.
func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://www.reddit.com/r/%s.json", reddit)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	r := new(response)
	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		return nil, err
	}

	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
		//fmt.Println(child.Data)
	}

	return items, nil
}

var desktop = os.Getenv("XDG_CURRENT_DESKTOP")

func Desktop() {
	fmt.Printf("%s", desktop)
}
