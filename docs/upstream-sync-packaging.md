# 上游同步与打包版维护流程

## 目录与原则

- 开发工程：`F:\Game_AUTO\MaaFGO\MDA`。`origin` 是 `dingi8528/MDA`，`upstream` 是 `1204244136/MDA`。
- 当前打包版：`F:\Game_AUTO\MDA-win-x86_64-v1.1.2`。它不是 Git 仓库；每次同步前先确认实际使用的打包目录。
- `F:\Game_AUTO\MaaFGO\MaaFgo_NEW` 只作为 MaaFramework 开发范式的参考，不从中复制 FGO 业务资源。
- 保留上游的非会员功能；移除会员验证、运行额度与计费、设备绑定、续订提醒、票券生成、额度显示，以及用于会员分发的 MirrorChyan 配置。洗词条的词条“配额”、游戏内付费商店和 VIP 道具属于游戏功能，不能按关键词误删。
- 会员相关改动的 Git 提交标题和正文使用英文。提交后不自动推送，推送由用户决定。

## 同步开发工程

1. 检查 `git status --short --branch`，确认工作区状态；执行 `git fetch upstream --prune`，比较 `HEAD...upstream/main` 的提交与文件差异。继续在当前维护分支上合并上游；若有冲突，逐项保留本地已移除会员功能的决定。
2. 查看新增的 Go、Pipeline、任务、语言包、文档和发布工作流。检查 `membership`、`RuntimeQuotaCheck`、`QuotaDisplay`、`ticket-generator`、运行额度、赞助/订阅和 MirrorChyan 相关引用。重点核对新任务入口是否又接入额度动作。不要仅凭 `quota` 或 `VIP` 字符串删除游戏功能。
3. 维持 `assets/interface.json` 指向 fork，且不含 `mirrorchyan_rid`、`mirrorchyan_multiplatform`；开发工程的 `version` 由工程自己管理，不拿打包版的本地版本覆盖它。
4. 运行 `go -C agent\go-service test ./...`、`npx @nekosu/maa-tools check`、`npm run check:theme`、`npm run format:check`，以及 `python tools\validate_schema.py --resource-dirs assets\resource --exclude-dirs assets\resource\announcement --interface-files assets\interface.json`。修改 Go 后重新构建 Agent。Windows 控制台若无法打印 Schema 校验符号，先设置 `$env:PYTHONIOENCODING='utf-8'`。
5. 检查最终差异和残余会员引用，使用英文 Conventional Commits 信息提交会员相关改动，不自动推送。

## 同步打包版

1. 对比开发工程的 `assets/resource`、`assets/tasks`、`assets/locales`、`README.md`、`LICENSE` 与打包版对应文件，列出新增、变更和需删除的旧会员文件。只复制有差异的文件，不整目录覆盖。
2. 在打包版 `backup/sync-no-membership-时间戳/` 备份将覆盖的文件与活动配置，记录复制前后的 SHA-256。不要复制或清空 `MDA.exe`、`maafw`、用户游戏配置、`cache`、`debug`，也不要删除打包版保留的 `resource/model/ocr` 模型。过时的会员设备缓存可以直接删除，避免在新备份中继续保存设备数据。
3. 按 `assets/resource → resource`、`assets/tasks → tasks`、`assets/locales → locales` 复制差异；复制已重新编译的 Windows `agent/go-service.exe`。该二进制用打包版显示版本设置 `main.Version`，并采用 `CGO_ENABLED=0`、`-trimpath` 构建。
4. 打包版 `interface.json` 从开发工程复制后，单独设置 `version` 为 `v999.0.0`，移除 `github`、`mirrorchyan_rid`、`mirrorchyan_multiplatform`，使该本地包不从项目更新来源自动更新。除非用户另行要求，后续同步继续保留这三个字段缺失及本地版本，不照搬开发工程版本。
5. 如上游重新带入会员文件，删除打包版里的对应旧文件及无用任务快照；用户配置只做必要的定点修改。现有 `config/mxu-MDA.json` 已移除 `settings.mirrorChyan` 和 `interfaceTaskSnapshot` 中的 `QuotaDisplay`，不要因为新版本同步而恢复。
6. 验证打包版 `interface.json` 与所有 Pipeline 的 Schema，核对复制文件哈希与源工程一致（`interface.json` 以本地版本和更新字段为例外），并确认活动目录没有会员校验、额度任务或付费分发引用。备份目录可保留旧文件用于回滚，不参与运行。

2026-09-24 的首次同步备份及哈希清单位于打包版 `backup/sync-no-membership-20260924-195440/`。
