// Config for the colors used in the tool
const (
        InfoColor    = "\033[1;34m%s\033[0m"
        NoticeColor  = "\033[1;36m%s\033[0m"
        WarningColor = "\033[1;33m%s\033[0m"
        ErrorColor   = "\033[1;31m%s\033[0m"
        DebugColor   = "\033[0;36m%s\033[0m"
)

// Config has been created
type Config struct {
        Insecure       bool `yaml:"insecure"`
        TimeoutRequest int  `yaml:"timeout_seconds"`
        Checks         []struct {
                URL          string  `yaml:"url"`
                StatusCode   *int    `yaml:"status_code"`
                Match        *string `yaml:"match"`
                ResponseTime *int    `yaml:"response_time"`
        } `yaml:"checks"`
}

// Config has been created
type CheckOutput struct {
        URL     string `json:"url"`
        Status  string `json:"available"`
        Elapsed string `json:"elapsed"`
}

type JsonOutput struct {
        Results []CheckOutput `json:"checks"`
}

func addEntry(results []CheckOutput, url string, active bool, elapsed time.Duration) []CheckOutput {
    // fmt.Printf("active %v", active)
        check := &CheckOutput{
                URL:     url,
                Status:  strconv.FormatBool(!active),
                Elapsed: elapsed.String(),
        }
        results = append(results, *check)
        return results
}
