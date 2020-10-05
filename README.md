toml config file, app config reader

### 查找顺序
- /tmp/config.toml


```go
import "github.com/wiloon/pingd-config"

config.GetString("group0.key0", "")
```

