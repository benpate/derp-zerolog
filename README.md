# derp-zerolog

This microscopic plugin for the [Derp error reporting library](https://github.com/benpate/derp) 
writes errors to your application's [ZeroLog](https://github.com/rs/zerolog) logs

### Usage
```go
// Import the plugin
import plugin "github.com/benpate/hannibal-sigs-w3ctest/derp-zerolog"

// Register the plugin with Derp
derp.Plugins.Add(plugin.Zerolog{})

// Yay! Now derp.Report will write errors to zerolog.
```