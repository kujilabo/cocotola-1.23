//go:build large

package gateway_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tcolgate/mp3"

	libdomain "github.com/kujilabo/cocotola-1.23/lib/domain"

	"github.com/kujilabo/cocotola-1.23/cocotola-synthesizer/gateway"
)

func length(t *testing.T, audioContent string) time.Duration {
	t.Helper()
	audioContentBytes, err := base64.StdEncoding.DecodeString(audioContent)
	assert.NoError(t, err)

	reader := bytes.NewReader(audioContentBytes)
	readCloser := io.NopCloser(reader)
	// r := io.NopCloser(buffer.New()(bytes))

	// streamer, format, err := mp3.Decode(readCloser)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer streamer.Close()
	// fmt.Println(format.SampleRate)
	// // assert.Equal(t, 0, format.SampleRate)

	d := mp3.NewDecoder(readCloser)
	var f mp3.Frame
	skipped := 0
	var x float64
	var y time.Duration
	for {
		if err := d.Decode(&f, &skipped); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		x = x + f.Duration().Seconds()
		y = y + f.Duration()
	}
	t.Fatalf("duration: %v", y.Milliseconds())
	return y
}
func Test_googleTTSClient_Synthesize_JA(t *testing.T) {
	t.Skip()
	ctx := context.Background()
	t.Parallel()
	httpClient := http.Client{}
	apiKey := os.Getenv("GOOGLE_TEXT_TO_SPEECH_API_KEY")
	c := gateway.NewGoogleTTSClient(&httpClient, apiKey)
	// audioContent, err := c.Synthesize(ctx, libdomain.Lang5JAJP, "ja-JP-Neural2-B", "私は母に髪を切ってもらった。") // FEMALE
	audioContent, err := c.Synthesize(ctx, libdomain.Lang5JAJP, "ja-JP-Neural2-C", "これらの服は洗濯される必要があるの？") // MALE
	assert.Nil(t, err)
	t.Log(audioContent)
	t.Log(length(t, audioContent))
	t.Fail()
}

func Test_googleTTSClient_Synthesize_EN(t *testing.T) {
	t.Skip()
	ctx := context.Background()
	t.Parallel()
	httpClient := http.Client{}
	apiKey := os.Getenv("GOOGLE_TEXT_TO_SPEECH_API_KEY")
	c := gateway.NewGoogleTTSClient(&httpClient, apiKey)
	audioContent, err := c.Synthesize(ctx, libdomain.Lang5ENUS, "en-US-Neural2-A", "Do these clothes need washing?") // MALE
	// audioContent, err := c.Synthesize(ctx, libdomain.Lang5ENUS, "en-US-Neural2-C", "I got my hair cut by my mon.") // FEMALE
	assert.Nil(t, err)
	t.Log(audioContent)
	t.Log(length(t, audioContent))
	t.Fail()
}

func Test_Audio(t *testing.T) {
	t.Skip()
	audioContent := "//NExAASGAXMAAjGABlAQAYPnAfD4gBMoCA0HzgPh8QAmOBA4D4YE4PlATHAgcE4YLg+oE3ggcE4IFwfgmfEBwuCEH5d6nTn///+XqOKBCD7xOfKAgr0KAfEwGBsXIGg//NExAoSabogAElGlKISYABUDnO2jfo2d/t7f/+8/hPCeJTCDAAccumIKCA7/EogOZB2Knh2AC+sLQWAYfcBO///6e/KEw/iAaH4WM26/9pEA0pjsiHlFrpUGveft4dH//NExBMU6cp0AMnSlJStrYNEhqv++v//////////57cyNoUMSZFZwktd5A3MjJyAMChlIAYbN0oKHUgJ5MEciAUKk8ItkBJh9/////9K/tieFbBIGnaQ0AxNzHQZ1hTz//NExBIW+dKMAMtelaBPlNCt/T61If/////////j6tWW8JCy2D1kvQhLJVSoeulG0q9VPEOVJfwbZcmmE1HIZBpuskHJIxQ5kOJw1qV8bCAcnOtWCUap/HVVVIHZ1onT//NExAkUMdKgAMnelDpC9Mou/0GF219ft/////////+//hspEyrxTSbqvTZBAtx3CDsORQVenmDkOxVq8yA+21alalCUjJqVzEYVlJlITd/7xIOHqf+vwdCgbM/bPKZ9//NExAsUqcKYANNelYjPwLesaCTIf60Pf///////r/+sSnfopwkbYIMEchDGxjE/BJKY7HhfBfiAKhUp0WZSyMz5MtD3ciuU8VWq1sVakVs16Q+fkubpadBCeBEsR4YS//NExAsUccKMANHWlM9NEzKAUiqzIkgj3MtYrb//Pf//TPjfW+zGGyYB6MQ8m4eWlo3C8DJoCeAMP4CoBJSRBqeO6jbRJseav1HnTUtStqBtf/+msZ/kX8OnRRde6AGw//NExAwTOa6EANJKlEFCS5y6OzyZif9EaOe8ttQz31xRHQoCBIgTWgRyEBFG5cxS1qVH0RucwkLOMAVQ8p5nQdX3Z24mU0bb/xooD9M15/4SIwNTSaUYl9emUxQSMssy//NExBIRWaKQAMpElM6JB6Q8PDhQNrFGmUZuAjAOSYFx+B8SpMs4VP////2ZpjodWOpmO7MlVYIhwOHGfp6Dcqr//CNKqmAgvu6/Ih9pGdQdp/bfNPSqd75iW+5ThAwA//NExB8RUaaYANYElBiO641M4ylkqYhnC6avKNc//9/////76JMkquYhL11OGGaUR/3GKvWsPZBHDMdPSGgLSLHh/G9PyuCxUgcjdBEIB4XTP1Ahk7AFaHm6qwV8+7/L//NExCwSMa6oAJYKlLnn//////uOIIILKQCCYfQcNIH2D4mUUQokKR5m+tgOwnDJOoyH0tiGPl9PrHr+BbqszXYpVxe2OAjoyJLC32lmmtQfDmGP//5f////+0ziIwRA//NExDYSSa6sAG4KleHDB4PCYIBxAMMQPHHB8QIE5b5V/WRAV8oJNWqVqaJrxO7YqS5R8zKzswFmH0xtQ08YELCNQ4u5fu5byzpOf//T/////38dGwnCAViWXPGoiiMF//NExD8RiaasAG5OlURhoNGBRYGq/pgBQK0m10jAgxtWlsqynHAgtYU/hZFovU6tTGTKrGHPUrp0uW//Uzb5///////1NdrnmIPgGgpBgLB01ReFkRY/NVX1qWYBBAFh//NExEsQ6a6gAG6UlKqmZRTAILiSMtjpqOSSwHDJg3OJo3Jcni8alMHOi4cvdRmYP/b/////1N5ymmAqMczVCMfTvyT0dTv////sKPW5fPA3KCridSPqfGaLAC/f1GtZ//NExFoRedagAJ0UmMsgc4DZlxflVam0sCLGzXPz1rLGHf/////////1AEtfUOtyX/////40TCgYDybv7wqWXpcsK1U85mxd7djiwkWZ1oopiBhJgVolCWLpiiXBhB2k//NExGcQcbKcAJ5KlARdlorVV////////////////////6/9Kul2kUgo6OJBMICj1e9qYTUYUwNhWNukQJRiU1Yg7qP0WuGnHHkwcg9B6a7iEaDYq1XOf/////////////NExHgSCyKQANNKvP//////////6OuUzrpEgCKLKMUOlAowcZWy+0TZaFxQsbNZLOMpDjLdi5xeqCi/xgwjNn5zEZzyRcmaNxqEYBQlD5IqOnqc9l////////////////NExIIR0yKIANHKvP/////00qzFKU7CDKr5EOzO61EWMLGDwGEW/9a2Yk8K8QksgNJGgO4gKSRw+m1TK9WH8bppFYyGkCm2p6QYr2NWtdb181xumd//dcX/z8f///6c//NExI0VAyJkANHKvN2///6Je3////////d0DzxgiYPCwGQwg4QFWMjCLLMzQhT//lJzNgtLQzBHTpb8YxoIxSzblxTmGT7t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hASIUAkCkC7JiQwyYKgsL7GEA8JGV1JDyg8MZXVj0b9///////f92/eeep9DlYfuxf2GeA7nF4mxeTMJSFlx9+UzmUu7SQWg//NExJEeaxKUANPUuRVn3bev/bsupY1+NWlmmFwzLdbs/Op7IJ4XY5Nv5JawXAcBL8sTFLnLHYLxxevjWr8eFYUMfcsV4elr9LJACkRI3MU9NYlUMxCkw13PCVROX/vV//NExGohme6kAJ5wmFq2XeZxO6s93N3KaHLNto0u////+lKBO0WDRoPq/SMwJgQFPYvJAl4YmYeHpHUAhQpJr6SQ0W9SY6S+hWynQEFi8ZpmBXclpTji7F0xipNxSNTK//NExDYdQgasAJSwmJCI2Lu70aylKz79HardlrMES14U2H52ZdL2KRXn//d3s//99mY847za5jcpYlBMCSHPuW+/cd////4BTFQEHP5KhTQV1waoslzXVJgg9T9ZxL8z//NExBQUMfq8AIQemFP9ApkVPNTMEWRo0Swza+YSnYldGrr50/rXX3A2XRjJSXFxrp6q2SFfX/+/jf/9PnGLVjKhkkUqh1bV3ir1KgRsZ5yc6hHHBS5rHOwAwTOrcqb///NExBYQ2f6wAJHWmFN+YNiRz6FwOgSlqLWuurbXzw6/+3NYbXulhqNpNr////4d+2HDcTa5a6TU1dW1jlfgBW0Rixv0qpspWghkcrdhKgaFaWshmuKKiq4GGNb/9S7S//NExCUSixqAANiKvbmcqHm///5NOn/////////////W9kdFmEFQPRc4scBBIe4KFvOq7vVWgetLs+4Be8rpZcYYmYi6AtT7YGT/N9v3//8MnkLmXNqZ/SBMNvD63//n//NExC0QkM6UANPScBqhhQPp/l3yhR1v///9uIpMWESkBqr/5lXnyzoZ1Szz+ywLBHFIPBUc1ORMKhGlvbu9Qw/UdSRffqSyORXLKZXPD6GbAp+xk/6X7Tl60l/8rdXL//NExD0c+facAM4YmLM51+mf38+ZpTv3KUYei0/M3g7MjVwsZQwhg6abfNpTrur6lt/7EwzaciiZfisZHf////S4zDH4nsDwW486AkYCkIAlz1SAAHNByrI2kvJ4OKmK//NExBwVCcawAIPKlZP0cyrMisH2Anj4CRLUFPlzPw0WV6yuX///p3ciiphogEggEzDXDoqKKjsKGbIJCCoRREULjKFMKvWDUSCmjrMwH2HKksfuVU7nWs4btpbuLS99//NExBoRkWa0AJYKleBhuGF590txKZcZm1LYhh3a28f3gb////no48ICIkLOkXZFS8opowVZr0d+quspANUaZo6hQ4B8RQT2w8HNhV6tZR4YNav35ahE22WUeERS3zF6//NExCYRqWKoAI4KlRfmhe+l1vnMcf////33OKxhnDBIPCw/xVMI4/FgtN7v1cvuUK5z3YUlVuZjJ5sAwJnhEyyq9a1rl+dpqeI2F9LWrZM9WgZRMe1Ulc7+8d6y//////NExDIR6VaUAM4ElP//yrMDASDip2WuIyx7RV9n/////IrlHSMjN3TsOMQzTcWFc1fkAtbsltboK0bzjQkSJYUOQyaIjucLN3SmT5JS5/z4XkgmuUuQqoPKhIHDDDY+//NExD0RyUZ4AMGGcKc/eobOZz/+74pV9WO45HDnui/ptKuYbEhM5BFAARQ4PKLEz9xhkXsYVkMpBpe8vdd8f/MjkFaWugkNZnzT7Q3T21EBIJzoDnHP778xofmGr28n//NExEgQsWaEABFMlOEfL253ynw2P5xx/7ulVt4MfLRRQSFEKWtFH5mFUlMdDDqEPir4mYRU4XHMeN//+PhNTy3/7Bi9qpFEh67K53PsMkAZ0WGcA/M1IILCxBCxK1Lq//NExFgQqPqMAHnMcXp6nrep7qdBRSlFgSCBDOGFX8KVtHe13mo9KqJhIFQ0v50CKKihH9YUSONCKEpXR/f/6P/XM50LoGJttJXXtY3UpgVBq+yJLv7JpvXUZIIIoopm//NExGgSAUaMAMtEcOOAc5T3Mx3X/88iXGi4fD53OMEA4CNSX4fWHA+Xtd////9Pa/jK9RdGeAFZBthNEs6B0Z4Ha3dMuDzf1lD8dxmX33Kg4QkYmcc9NEc7MzLMzM5k//NExHMRYVacAMNKlKowc6cZBERgBBEcxxCMzAmOTs1eYp031NYwPKTUpmGmB4cIhEPf///6C1zxIpwoGnlylfWoogG6H1OF7mgKU6/MB5fWZJfRHsbOpaKyVEaJa4Vx//NExIAX2bqkAKNYlAAnD6/rglf+YGrVyg4TgAQW9LYoI5tdcpf/s0r/RwzMu////9X21OowLBOgdGCcSkRHM3DxitYych2p9vUnvnkA97CiAoBYLmJBiGVfxUvtWohi//NExHMRua6wAJNQlPbszoCsBe3n9D5RQRrKOyCf/////VX//LLB0jKwaSeda3LpkcAAK+MwpIuG5QDoCdT1Y40NOeEKFVSqJOBnAXeTJfzkb6rDmfCpKG5ZWtxh4HM+//NExH8QWW6gAJnQlKcuTdI9UggfK9XgVTpB9RSIz0ThxZgIOXqs+GK2MEPgxfDv38YunPFr/65EHQbTuOtbp33tmzoQd///9qiAs4gK3lAA5wPhZf/f9zeIypLe4r3C//NExJAhIcqYANSylNpVCde9tWYh6ZaboJi3dVAkhvMyVQgMDYgvOYmjCewgEFyOomSlE8KXKp5qzw6RbDBNjYhwhEADAnEkUCOGsQ0SgC2juSeaGpOEKtqlFEay1Jzh//NExF4gsxagAMykuMHULaXEVmaBOjMO7XSOI///////6drrXdCvTRUgrUdQZmQQPH31uHwC7I02SUUACUAMYvoM6x/KLe5r9ZmPU9qUVnDV0/dXloW5zzFxAgzev/xd//NExC4XAcKwAGtelO1+fp82svjQJFGXKdTPo88CeLaf/2zWv+aRs1xjEGSMoDAEBK////9RWPDaFBX8t6mXQAKAU5/56llUndmQgYWYRFUcvlKnlKrbzOVSlawChs08//NExCUR4VqYAMFMlLVVf///tr1VU5GCRuscSWdEp4735Y7gr////8seiWkrDWJV1XNnKMklEaSeXUZaOIwBjLR/AAFkymQjQOr0qdL1pzqmpUQG9m073uX5Vyr581v///NExDARQsJEAHjEuP+n////////83uhatHkh5wSKBORj9z/+pQtaCpcoJFECPWF289o4o251k1yMwjbnSCBASTnKAjA2aFDlMmj2EPSOc5z2D072yZMmUQy7JpxjR/n//NExD4c0w5kAEpMuGjP/nj972Lvbvve3d7+0eI/vWjIIEJ38yIe2e95gDTMxz04zmEJJpxN+I/cnuR8u/rPGWxDFd9qXaNpDLJmNw2FkjmoBVjxxaiRDBR28pM8Y3b3//NExB0V8c6cAM4ElG/uy+//c62HbudBSuXI6SphG3/tMWavJ9Pn9auCIIUDOEAGQjIxP+QnoQnDh//1+v7N2nSuC6IPqPiB//1NkIcVptS1hxzsrcwlUTBPUMqsovxx//NExBgXYbaoAM4KlGJI8fpJXEK97OAEvGa1re3DBJ4Mr1KWZfmmrDylN6L0b3bkU1YmGh0wcDgQCKsRDF9nrVCutx7BUPFw4WTp////1tjE1KX//9JEV7t+SKlDXozh//NExA0SmVKsAM4ElI2Q5sTry3T1S/Llq6xm5YtyhhSpZPli/rYGbZ4446//o///+puoYVVxQCROGq3VPUInmg6soeUt2/////1q5fls1LpS4LatKKohx2G+E2GQStSB//NExBURqUqMAMvElHx9tsCNCfT6LqJqtHKhJfR6ThP5zkeq2Ryr///////+bqFf/54FWrdDhYO0f////2y1bmYgBf1QpxJUSnS7LKFBhCWAfIN9FgEpDPDiWikgITET//NExCESieo8AHwEmBcpEhQJcK5Mk6amxisydl///////t///6LVqPahjiTHga////1bZ6VGqttSmmqlA2jE6RCPwILnGORoT75Ii5uVzvAF4L8FIhrEkTjTjMnHrPI8//NExCkRefZMAEvKmKUpumv/8a+Kf////////kJIroAgpoIKHmf////V+UOQrYWwVVMcszR1Xp+G5sEmnrqVnvSrtAWOgAcwimCEosYngGoQCliEgwANTAMxQLnAOOQQ//NExDYeqwJwAMzKuUBg5UPOQQfBfJY3LxoYHqrNv///fsczHGAcSFFQjey1ej1chzBzRhAABx2YUMNMIjz01d////////6EVGUPizCgmeBnOvy/9lgDD1NfjjXjAyw4//NExA4UycKkANYOlHNtboHFKwrWitbsSf2elNajtAKjfQ5g6bYoBnY3nUx7///9////qiDQxlJisPNex6ffzFoqnk2ow4iWO/////zI0UF2CtX3FakHTY6QwD5gcotG//NExA0Vyb6sAJ5UlNaQ7Dbrf59xhhqVrL+PtAtbdKqiOnMSt3raMiNEgr59lc/2k/5n39TtCo+NYxAZC+LEqxcOGMrEavydUOnFkOzS+DJalv////6UKv/lqBGGDeIZ//NExAgREV6wAHvQlAAAHnhekosxgq5yrhTKrG/k/XH06+BlOGSOow2h/pWNi7pymv6/4n+nb+Sr03BqsUwDQharUPuJ3bvU+qry0DY2HuITwXzAuRcQeEGYLBCj/L/f//NExBYR6W6sAJPGlP4bPf03qm48JOkrZ4MlCDhyIyIhBxf/fm4tCiVz8wARXMWLxFBABHOiDP///////29V6nFAmLUhIQep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7h6OpWJz1UsG92/0su985cxPcpGpBUBpovp8S6P/uqDkNHpIt////vzuSZK1/uE1//NExI0QKWK0AJlYlB884tNBXU4/LEggcFAahL9BcBQARY6yHPzv//1WtSAFhIC46ynFUaG/4aZrZxWHUaDUSB9lgr/509d////x4aSpKjzSz0h1Csbte2vo4UYHMGaW//NExJ8SGW6kAMiYlJrZi2514oc3fF7oRDj+KNn3ikTWPDZ5/ilNf/cB5rH+Hms3wnFQ8Q800sAuADwcBoMh+GOWx7PsKAGgeCIFsuFwBQAUEwnAkCwBQJSeYK5ckzFe//NExKkSmVp4ANHQlHuef9pkweDwkkBURYiyR2Iye//9X/ar/P/89/V0bbvQwkgnSCH/xxlCByTPIabslSbjYyN2po/x6wP7DPSzBr/NFZ//ePr/0l/8Nkh/GWNHzMqM//NExLEguup8ANPUuB6xzhH2uN4QCOBgIYP+whtGjOa5TQJ3c21ttopTvUXcRsIis4kb6QmJK4npmejdSi7+rG//WGGP4551LOd2STkG01FSOw/E6DAnR///77JNzAKa//NExIEhEeaYANPwmDayBhX//+rRqZ36YRGHUHFKOwM4BjCxInpgb9STdND0EvPdZma0TArjQFBRdyUAY6AGPASYBLyIQID0FWKxRiDyPzy7Va27UPwy7UDJUOGoK3sp//NExE8cue6sAMxwmIYisWq/jb1lrX/l/59x1lqxd1GspVST8at2M79UO/////6HVP/hrL1hq09BjhuOWsBRAETEOnen29ft+P/4UGLixmklLjEXR2AqjCS9ZnJW0+av//NExC8RUXK8AHqelLeYtDebWXEJMsr3VtRmWfsEoaCgGPJq/DDFYMeTT3cH/OGxYNjmILgdF1RzW///////77wqJWhScCoATOulaGMt8Y/qoUpweqGWVYCUb+su8qgG//NExDwR+WqoAMHSlIEy5Jbr9EuD4IAMP9HChBTGQCW////////hkir//////v79/vxv/8/7Zr/uuUYvnIROpGg5Os3vhmKow2iCaYIGo1IrAGUEIIIMgFwfLBJDT6St//NExEcSMmKsABBMuJ+h///////4v///7/5rji7u/mrbl7jEMYKnFBwIYODA/JhA7MMPg4WcadSkmkDSpWg6HWeLDSBcch5gqlOjSWLBw0I7C3QvclX///////l///////NExFERknqsAABQuP9+ZUav//1l55+v1ka+titM0Kg2Jl1DxOjaSJRf6gKqU5KkbLDRO+KSPZjpQlXkq2skDFzSyFeneq+2ad636f/////5fkP///+Km72ffVO5rvVX//NExF0SMl6oAABSuaVXUJZ0kD+kwfCAgGxwuakpMlG9RLmLnXn2m9E41cmbKrvOIprsXhzmNRXCQeFRMHar3C26hZveABABDRe8IZbn5OBEUfAD5///x7gsRhAcIjA1//NExGcSYmqoAABWuB/+QQIYXcyFXXfTqViGsHRIBhYWdRFiobZStapS1qaUoqwqWPTy+JX/3xhS6b1S81UFGvZ4u2B6ZoVkH6ZTP6kDTWmTi7miTJuUBwG8ao4ZupLJ//NExHARMjKsAGBKuOXsHXv+blYvTHykakRWActsozbyqLiWkwkMV6GxVf/zYQeK1RIHmbdEiAeE/WRggVtRENZ0gbKWmxi6xrhaGN5NNMZQW5UcCtfnf9/y1/I9uICQ//NExH4SIWqwAMQSlE7ooGlJNqdHLr+8X8Ghom6EaPGYM9JxoBMuyYTAXD/7aMmva3u4/fe5XXY+3UaqRTzqakpWK3a/b9X8c/09PN9n2GZaxGA+J6UEYnPyvy/Qc5Ut//NExIgRwZ64AHwQlFkXq0rrH4POyjYPIJik5TCUCIsa3t6Cb3Nfip/m9RAeRd94ipd+H+kObdorTZ/c5z636v/fzaudLFssfQ8GC5OHmlnEOs25kSvODxayy3af3UXx//NExJQRKbK0AG4UlPAgao3DhConlGANMIw0us2UDTt/95oXMPbcWbKscwr8vzE4p70Zfr6nP+vp//35GS0KgHgVcYAFAvnoKhvK8wZOiEYtMxyFkqdT//////Wq8UwQ//NExKITadKsAJ4alOWcTEiECqkwA+h52X6goQLJ9WMqBDhLqkwz8E4rVjczYCGWh0czfi/O/9//+Z6/9/fyyGIF2JhMxgBgNZZkEALL8m5QQuMwoxQexUQAsoI57/////NExKcUCbqsAJ4UlP/p//0n4327lxW9iv6bKYuKOMGUslDIpNl8rRHmaavBJvNOlGJAquGtPfVxryD//fP/1N/q/7eYsqTBGSVcyFmAkhuRckCXRzMlnzQczVhXmLqL//NExKkWAbaoAJZUlIYkjUdcxmCMWep3///1743RN+oU5CtQaimg4/ANki93bgDYYzlyC3Ix3MIeB9YzXh8dSk1rtWP8//3/////////H/F+8lRJu8CAnPNx7BOPOem7//NExKQYuc6kAM5alODQut6BqXOY8kGzWINEonYgh/////5V8tlqtAUxNAzEQAnl4yTAphbTZ0w/MtG5gfAB4tpeKIEaDbItmqZK+g/////6nmzCgDAckCo3BwEyjgiH//NExJQV8cKwAJ4WlDsYOHTjBQNnHhoI4kseQLtL4eml8hvMxwstQYEDalqLAWPkgfWOsiyy6TYXABsRicEBAqJuVUibLz7/////+p5IaxEB4BESWGgiOynnG3PHjmMG//NExI8T4bq0AIROlISKIYOqp5d9/qryk+PukiOsA9E9Z4HwMZ7Z4eMFD6gTZOjnVQtgA/MW6vfM3/1////////axQdH2LgsJiSJBU5DlPY50NKmjjEjTFLiiXa11f87//NExJISCa68AISOlcpnNMvEpYCIxBKtCTBEMt/Gcaey9SxlLlYBwgVIn3J09YVYrWP5z////////9M084bGhVRwHxkXFUGTjTGNYnY000fNG094spmaz61aylMRxGgn//NExJwRibK0AIvOlCEnEvFipd5t5C2sjsU7kSJ0SzAZYFKQvSFfi5NRmW8//////////zaoVQMIh0CuMcbEnKOVA8PoKai2Nyr/+2akqiRFRKIgBKOgoMUZiu1FF9IX//NExKgSaaqwAF4OlRmKppIrMBFUHeiFyU6A1dsCyr8u//////////+ZQ6URHkAYPOMDwsLcFLir5DYRx03z1vOtlsxaZYBVGEDADJPyaKpXTwYjMpV2DBAaQJ4T8NE7//NExLER4YqgAGYKlZNvYOv/////////wwFUYWJIYU8rGo8xQJzLHs+VRZ/////2KW00iRytNSOWAQIKSNx9AVGID4ziVFyTZ0EoCYfmj0XNSmL/5/9Df6//70MwE7Bg//NExLwRuWZ4AE4KlSPag6WBURFhKQZ5b+tywk///+HYi8OKEZZkZEZforPqiqi//lRVRUVFQGDQ7P/6oqo7OxioioqKgMECUjt/uUyggSkdnb/1//cpgoYGDCHI7f/r//NExMgRcZZEAFvElP/ZymCggQkDz+VCoYpMQU1FMy4xMDCqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqTEFNRTMu//NExNURWV4QADLElDEwMKqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqTEFNRTMu//NExOISItVgAAiEuDEwMKqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqTEFNRTMu//NExKwAAANIAAAAADEwMKqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqTEFNRTMu//NExKwAAANIAAAAADEwMKqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqTEFNRTMu//NExKwAAANIAAAAADEwMKqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqTEFNRTMu//NExKwAAANIAAAAADEwMKqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqTEFNRTMu//NExKwAAANIAAAAADEwMKqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqTEFNRTMu//NExKwAAANIAAAAADEwMKqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqTEFNRTMu//NExKwAAANIAAAAADEwMKqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq//NExKwAAANIAAAAAKqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq//NExKwAAANIAAAAAKqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq"

	audioContentBytes, err := base64.StdEncoding.DecodeString(audioContent)
	assert.NoError(t, err)

	reader := bytes.NewReader(audioContentBytes)
	readCloser := io.NopCloser(reader)
	// r := io.NopCloser(buffer.New()(bytes))

	// streamer, format, err := mp3.Decode(readCloser)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer streamer.Close()
	// fmt.Println(format.SampleRate)
	// // assert.Equal(t, 0, format.SampleRate)

	d := mp3.NewDecoder(readCloser)
	var f mp3.Frame
	skipped := 0
	var x float64
	var y time.Duration
	for {
		if err := d.Decode(&f, &skipped); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		x = x + f.Duration().Seconds()
		y = y + f.Duration()
	}
	t.Fatalf("duration: %v", y.Milliseconds())
}
