# cosmos-address-converter

코스모스(Cosmos) 블록체인 계열의 AccAddress를 사용하여 해당 Address의 ValAddress와 ConsAddress로 변환해주는 Go 기반 CLI 프로그램입니다.

---

## 프로젝트 구조

```
.
├── config/                # 설정 및 환경 관련 코드
├── infrastructure/
│   └── grpc/              # gRPC 클라이언트 및 에러 처리
│   └── modules/
│       └── staking/       # Staking 모듈 연동
├── internal/
│   └── address/
│       ├── domain/        # 주소 구조체 및 변환 옵션
│       └── service/       # 주소 변환 로직
├── main.go                # CLI 엔트리포인트
├── config.toml.example    # 예시 설정 파일
├── go.mod / go.sum        # Go 모듈 및 의존성
└── README.md
```

---

## 설치 및 실행

### 1. 설치

```bash
git clone https://github.com/mingi3442/cosmos-address-converter.git
cd cosmos-address-converter
go build -o cosmos-address-converter
```

### 2. 설정

`config.toml.example` 참고하여 프로젝트 루트에 `config.toml` 파일을 생성합니다.

```toml
grpc_url = "localhost:9090"
account_address = "cosmos1abcd..."
```

### 3. 실행

```bash
./cosmos-address-converter
```

실행 예시:

```
Account address:             cosmos1abcd...
Validator address:           cosmosvaloper1abcd...
Consensus address:           cosmosvalcons1abcd...
```
