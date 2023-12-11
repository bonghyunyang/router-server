# Monitoring Router Server

모니터링 프로젝트 라우터 서버

## Getting Started 

Golang 기반의 RESTful API 서버입니다. 

- Go Version 
    - GO 1.18
    - Gin FrameWork
    - SQLite
- Gin FrameWork
- GORM
- React CORE UI

- VSCODE IDE (Go Extension 사용)

### 설정

*필요 라이브러리 go.mod 파일 참조

- 해당 프로젝트 다운로드 후 go mod tidy로 필요 라이브러리 설치 진행
- go run main.go 또는 go run .으로 프로젝트 실행
- 기본 생성되는 DB 내 기본 admin 생성되어 있음

### 실행파일 생성 (ref:https://go.dev/doc/install/source#environment)

*Mac 기준으로 작성(mingw-w64; https://formulae.brew.sh/formula/mingw-w64)

- MacOS
    - `CGO_ENABLED=1 go build -ldflags "-w" -a -o router-server_MacOs .` 
- Windows(32Bit)
    - `CGO_ENABLED=1 GOOS="windows" GOARCH="386" CC="i686-w64-mingw32-gcc" go build -ldflags "-w" -a -o router-server_win32.exe .`  
- Windows(64Bit)
    - `CGO_ENABLED=1 GOOS="windows" GOARCH="amd64" CC="x86_64-w64-mingw32-gcc" go build -ldflags "-w" -a -o router-server_win64.exe .`

### 특이사항

- 현재 커밋된 코드는 최종 코드의 일부분으로, 전체적인 구조 세팅에 참고가 될 수 있는 코드입니다. 
