# Linkmate — simple self-hosted url shortener

### Migrations

```bash
goose -s create -dir migrations <migration_name> sql
```

## Roadmap

### Business features

- [x] Short link CRUD
- [x] Tags
- [ ] QR Codes
- [ ] Limit access
  - [ ] Time to live
  - [ ] Max visit count
- [ ] Import/Export links
- [x] Analytics
  - [x] Visit counts
  - [x] Top links
  - [x] Locations
  - [x] Devices (Browser, OS, Size)
- [ ] Api keys
- [ ] User update
- [ ] Multi domains
- [ ] Pages
  Simple pages with social links

### Technical features

- [ ] Add request data validation
  - https://github.com/go-playground/validator
  - https://echo.labstack.com/docs/request#validate-data
- [ ] Improve error messages
- [ ] Add logging to Loki
- [ ] Add tests
  https://testcontainers.com/

## Credits

- Web framework by [Echo](https://echo.labstack.com/)
- Migrations by [Goose](https://github.com/pressly/goose)
- GeoIP data by maxmind.com

## Similar

- https://github.com/ccbikai/sink
- https://github.com/shlinkio/shlink
