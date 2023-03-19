# protoc-gen-go-temporal

a protoc plugin for generating temporal clients and workers from protobuf schemas

inspired by [github.com/cretz/temporal-sdk-go-advanced](https://github.com/cretz/temporal-sdk-go-advanced)

## Getting Started
1. Install [buf](https://docs.buf.build/installation)

2. Initialize buf repository
```shell
mkdir proto && cd proto && buf init
```

3. Add dependency to `buf.yaml`
```yaml
version: v1
deps:
  - buf.build/cludden/protoc-gen-go-temporal
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

5. Define your service  
<small><b><i>note:</i></b> see [example](./example/) and [test](./test/) for more details on generated code and usage</small>

```protobuf
syntax="proto3";

// buf:lint:ignore PACKAGE_DIRECTORY_MATCH
package mycompany.foo.v1;

import "google/protobuf/duration.proto";
import "google/protobuf/empty.proto";
import "temporal/v1/temporal.proto";

service Foo {
  option (temporal.v1.service) = {
    task_queue: "foo-v1"
  };

  // Workflows
  // =================================================================

  // LockAccount provides a mutex for an account
  rpc LockAccount(LockAccountRequest) returns (google.protobuf.Empty) {
    option (temporal.v1.workflow) = {
      default_options {
        id_fields        : 'account'
        id_prefix        : 'lock'
        id_reuse_policy  : ALLOW_DUPLICATE
        execution_timeout: { seconds: 3600 }
      }
      signal: { ref: 'AcquireLease', start: true }
      signal: { ref: 'RenewLease' }
      signal: { ref: 'RevokeLease' }
    };
  }

  // Transfer amount from src account to dest account
  rpc Transfer(TransferRequest) returns (TransferResponse) {
    option (temporal.v1.workflow) = {
      default_options {
        id_fields        : 'src,dest,amount,uuid()'
        id_prefix        : 'transfer'
        id_reuse_policy  : ALLOW_DUPLICATE_FAILED_ONLY
        execution_timeout: { seconds: 3600 }
      }
      signal: { ref: 'LeaseAcquired' }
    };
  }

  // Signals
  // =================================================================

  // AcquireLease enqueues a lease on the given account
  rpc AcquireLease(AcquireLeaseSignal) returns (google.protobuf.Empty) {
    option (temporal.v1.signal) = {};
  }

  // LeaseAcquired enqueues a lease on the given account
  rpc LeaseAcquired(LeaseAcquiredSignal) returns (google.protobuf.Empty) {
    option (temporal.v1.signal) = {};
  }

  // RenewLease enqueues a lease on the given account
  rpc RenewLease(RenewLeaseSignal) returns (google.protobuf.Empty) {
    option (temporal.v1.signal) = {};
  }

  // RevokeLease enqueues a lease on the given account
  rpc RevokeLease(RevokeLeaseSignal) returns (google.protobuf.Empty) {
    option (temporal.v1.signal) = {};
  }

  // Activities
  // =================================================================

  // AcquireLease enqueues a lease request for a given account 
  rpc AcquireLease(AcquireLeaseRequest) returns (google.protobuf.Empty) {
    option (temporal.v1.activity) = {};
  }

  // Deposit amount into an account
  rpc Deposit(DepositRequest) returns (DepositResponse) {
    option (temporal.v1.activity) = {
      default_options {
        retry_policy {
          max_attempts: 5
        }
        schedule_to_close_timeout: { seconds: 120 }
      }
    };
  }

  // Withdraw amount from an account
  rpc Withdraw(WithdrawRequest) returns (WithdrawResponse) {
    option (temporal.v1.activity) = {
      default_options {
        retry_policy {
          max_attempts: 5
        }
        schedule_to_close_timeout: { seconds: 120 }
      }
    };
  }
}

// ...
```

6. Generate temporal worker and client types, methods, interfaces, and functions
```shell
buf generate
```

7. Implement your activities, workflows, and worker ([see tests for an example](./test/simple/))

## License
Licensed under the [MIT License](LICENSE.md)  
Copyright for portions of project cludden/protoc-gen-go-temporal are held by Chad Retz, 2021 as part of project cretz/temporal-sdk-go-advanced. All other copyright for project cludden/protoc-gen-go-temporal are held by Chris Ludden, 2023.