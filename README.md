# protoc-gen-go-temporal

This is a fork of [github.com/cretz/temporal-sdk-go-advanced](https://github.com/cretz/temporal-sdk-go-advanced), slightly modified to make it easier to integrate with [buf].

For general usage, see [https://github.com/cretz/temporal-sdk-go-advanced/tree/main/temporalproto]

## Getting Started
1. Install [buf]
2. Initialize buf repository
```shell
mkdir proto && cd proto && buf init
```
3. Add dependency to `buf.yaml`
```yaml
version: v1
deps:
  - buf.build/cludden/protoc-gen-go-temporal
breaking:
  use:
    - FILE
lint:
  allow_comment_ignores: true
  use:
    - BASIC
```
4. Add plugin to `buf.gen.yaml` and exclude it from managed mode go prefix
```yaml
version: v1
managed:
  enabled: true
  go_package_prefix:
    default: github.com/foo/bar/gen
    except:
      - buf.build/cludden/protoc-gen-go-temporal
plugins:
  - plugin: go
    out: gen
    opt: paths=source_relative
  - plugin: go_temporal
    out: gen
    opt: paths=source_relative
    strategy: all
```

## License
Licensed under the [MIT License](LICENSE.md)  
Copyright for portions of project cludden/protoc-gen-go-temporal are held by Chad Retz, 2021 as part of project cretz/temporal-sdk-go-advanced. All other copyright for project cludden/protoc-gen-go-temporal are held by Chris Ludden, 2023.