# Auth Service

One module owns User end-to-end:

```text
internal/
├── app/
│   └── user/
│       ├── models/
│       │   └── user_model.go
│       ├── store.go / store_gorm.go
│       ├── profile.go / account.go
│       └── transport/
├── server/grpc.go
├── config/
└── infrastructure/
```

See [../../docs/architecture.md](../../docs/architecture.md).
