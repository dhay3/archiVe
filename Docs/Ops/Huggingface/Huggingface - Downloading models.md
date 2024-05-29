# Huggingface - Downloading models

可以直接使用 `git clone git@hf.co:<MODEL ID>` 下载 model，例如 `git clone git@hf.co:bigscience/bloom`

但是一些代理会重置 22 端口的流量，会导致 models 下载失败。同时 hf 不像 github 支持 443 转发 ssh 流量，为了解决这个问题，可以使用如下方式下载 models

```
from huggingface_hub import hf_hub_download
import joblib

REPO_ID = "YOUR_REPO_ID"
FILENAME = "sklearn_model.joblib"

model = joblib.load(
    hf_hub_download(repo_id=REPO_ID, filename=FILENAME)
)
```

**references**

[^1]:https://huggingface.co/docs/hub/models-downloading
