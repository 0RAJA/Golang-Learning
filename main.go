package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {

	url := "http://222.24.63.106/api/submission"
	method := "POST"

	payload := strings.NewReader(`{"problem_id":9,"language":"C++","code":"#include<iostream>\n#include<algorithm>\n#include<cstring>\n#define x first\n#define y second\nusing namespace std;\nconst int N = 1e6+10;\ntypedef pair<int,string> PSI;\n\nint n,t,nothing;\nstring s;\nPSI a[N];\n\nint main()\n{\n\tcin >> n;\n\tfor(int i = 0;i < n;i++)\n\t\tcin >> a[i].y >> a[i].x;\n\tsort(a,a+n);\n\tfor(int i = 0;i < n;i++)\n\t       cout << a[i].y << \" \";\n\t\n\treturn 0;\n}","contest_id":"1"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("X-CSRFToken", "YuCpJhaGYv9QCiL9mjheQV6Qe59d0zbTGIfWveSupvRCYGMfKILyY340LjmySVIv")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "csrftoken=YuCpJhaGYv9QCiL9mjheQV6Qe59d0zbTGIfWveSupvRCYGMfKILyY340LjmySVIv; sessionid=s87f3ak6zr5h712ougk2nz19y5n9ku24")
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rand.Seed(time.Now().UnixNano())
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			res, err := client.Do(req)
			if err != nil {
				//fmt.Println(err.Error())
			} else {
				r, _ := ioutil.ReadAll(res.Body)
				fmt.Println(string(r))
			}
		}()
	}
	wg.Wait()
}
