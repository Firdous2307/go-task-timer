## Persistent Storage with SQLite

Go-Task-Timer now implements SQLite for persistent task storage, enhancing the application's functionality and user experience. This addition brings several benefits:

1. **Data Persistence**: Tasks are now stored in a SQLite database, allowing users to retain their task history even after closing the application.

2. **Improved Scalability**: SQLite provides efficient data management for a growing number of tasks, outperforming in-memory storage for larger datasets.

3. **Data Integrity**: SQLite ensures ACID (Atomicity, Consistency, Isolation, Durability) compliance, maintaining data integrity even in case of unexpected shutdowns.

4. **Lightweight and Self-contained**: SQLite is a serverless, zero-configuration database engine that's perfect for applications like Go-Task-Timer.


The implementation uses the `database/sql` package along with the `modernc.org/sqlite` driver, providing a robust and Go-idiomatic way to interact with the SQLite database.

## SQLite Implementation Considerations

While SQLite offers numerous benefits for Go-Task-Timer, there are some important considerations and potential challenges to be aware of,

## No CGO Dependency
The modernc.org/sqlite driver does not require CGO, which simplifies the build process and deployment:

**No CGO_ENABLED Flag**: You don’t need to set CGO_ENABLED=1, making the build process more straightforward.

**No C Compiler Requirement**: There is no need for a C compiler, reducing dependencies in the build environment.

**Simplified Cross-compilation**: Cross-compilation is straightforward since there are no CGO-related complexities.

### Addressing Potential Issues

With the `modernc.org/sqlite driver`, you won’t face the CGO-related errors commonly encountered with `github.com/mattn/go-sqlite3`. 

Simply ensure you have the modernc.org/sqlite driver installed:

```go
$ go get modernc.org/sqlite
```

### Choosing the Right Approach

I choose to use `modernc.org/sqlite` because it is a CGO-free SQLite driver for easier builds, deployments, and cross-compilation without compromising on performance or features.

And this project's aim currently is to make it simple, it might change later but let's see.

[Learn about Go-Task-Timer's features](features.md)