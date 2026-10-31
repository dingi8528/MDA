# Android CI 测试包

工作流 `.github/workflows/android.yml` 提供 meta → agents → apk 三阶段构建，产出国际服、Shizuku/root、后台 1280×720 的 arm64 测试 APK。仅上传 Actions artifact，资源配置不携带自动更新来源。

在 Actions 中手动运行 **build-android-test-apk**，下载 `MDA-android-arm64-test-` artifact。PR 构建使用独立的 `com.aliothmoon.maafw.mda.pr` 包名；手动构建和本地开发使用 `com.aliothmoon.maafw.mda.test`。产物同时包含 SHA256 和 JSON 构建信息，artifact 保留 14 天。

## 固定测试签名

手动构建需要在仓库 Actions 中配置以下 Secrets：`ANDROID_TEST_KEYSTORE_BASE64`、`ANDROID_TEST_STORE_PASSWORD`、`ANDROID_TEST_KEY_ALIAS`、`ANDROID_TEST_KEY_PASSWORD`，以及 Variable `ANDROID_TEST_CERT_SHA256`。缺少任一项或证书指纹不符时构建失败。

沿用已有 keystore 可持续覆盖安装；更换证书前先导出日志与数据。PR 使用 runner 临时 debug 证书，不提供固定测试私钥，不保证不同 PR 产物之间可覆盖安装。私钥、密码和 Base64 不得进入仓库或日志。

## 版本与验证

`Android/ci-versions.json` 锁定外壳、框架、NDK 和 SDK。CI 独立检出外壳到 `.cache/android-shell`。版本来自最近的 `vX.Y.Z` tag，生成 `X.Y.Z-android.ci.<run>.<attempt>+<commit>`，统一写入 APK、PI 和 Go Agent；构建号为 `100000 + run_number × 100 + run_attempt`。

Agent 阶段执行 Go vet/test、Android 资源与 CI Python 回归，再交叉编译检查 ELF 和版本。APK 阶段验证包名、版本、arm64 库、OCR 模型、Agent build ID、签名和 16KB ZIP 对齐。

CI 检查不代替真机验证。安装后先运行「Agent 自检」，再测试「启动国际服」与「进入大厅」。实体按键任务的限制见 [Android 测试版](README.md)。故障时导出日志 ZIP，不清除游戏数据。

维护者有本地现场图时，可设置 `MDA_MAA_LIB`、`MDA_BATTLE_SCREENSHOT` 和 `MDA_SHOP_SCREENSHOT`，从仓库根目录运行 `go -C agent/go-service test . -run '^TestAndroidPipelineReplay$' -count=1`，执行离线图像回放。
