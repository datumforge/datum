# Storage utils

The goal with this package was a basic wrapper for manipulating storage backends without needing to implement every client type and operation in our core code base. Today we have a pretty decent pattern around instantiating "Managers" in the server opts -> handlers, as well as the ability to add the dependencies in `entc` to be accessed within our `ent` storage layer.

All of the interfaces in this package will be for _remote_ object storage or basic file system storage / manipulation - not stored within our actual schemas. This means additional work will need to be done to create pointers to where the objects are so they can be referenced alongside our organization ID's or user ID's with the paths.

## Usage

With S3:

```go
storage, _ := s3.NewStorage(s3.Config{
    AccessKeyID:     accessKeyID,
    SecretAccessKey: secretAccessKey,
    Region:          region,
    Bucket:          bucket,
})

// Saving a file named mitb
storage.Save(ctx, strings.NewReader("HELLOOOOOOOOOOO"), "mitb")

// Deleting the new file
storage.Delete(ctx, "mitb")
```

With a file system:
```go
dir := os.TempDir()
storage := fs.NewStorage(fs.Config{Root: dir})

// Saving a file named test
storage.Save(ctx, strings.NewReader("(╯°□°）╯︵ ┻━┻"), "test")

// Deleting the new file
storage.Delete(ctx, "test")
```

## TODO

Aside from, like, supporting every storage interface available (super reasonable to do) the main TO-DO's related to some solid usage with our ent `template` schema and the main ecosystem of tools we have:

- "Fetch" functionality (open+write to destination)
- Additional abstraction with spf13/afero

### Nice to have's

- Benchmarks
- Throughput tests
- Supported content types / file types filters
- Uploader utilities + wrappers
- Ent document ID + object ID mappers
- Compression?
- Encryption