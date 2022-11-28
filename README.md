# Pulumi docker-buildkit golang sdk

This is the golang sdk for https://github.com/MaterializeInc/pulumi-docker-buildkit
The above provider does not supply a golang sdk so the following code was used
to generate it:

```diff
+++ b/cmd/pulumi-sdkgen-docker-buildkit/main.go
@@ -23,6 +23,7 @@ import (
        "github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"

        "github.com/pkg/errors"
+       gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
        jsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
        pygen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
        "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
@@ -182,6 +183,7 @@ func run(version string) error {
                        },
                },
                Language: map[string]schema.RawMessage{
+                       "golang": schema.RawMessage("{}"),
                        "python": schema.RawMessage("{}"),
                        "nodejs": schema.RawMessage(`{
                                "packageName": "@materializeinc/pulumi-docker-buildkit",
@@ -210,12 +212,20 @@ func run(version string) error {

        jsFiles, err := jsgen.GeneratePackage(toolDescription, ppkg, extraFiles)
        if err != nil {
-               return fmt.Errorf("generating python package: %v", err)
+               return fmt.Errorf("generating javascript package: %v", err)
        }
        if err := writeFiles(filepath.Join("sdk", "nodejs"), jsFiles); err != nil {
                return err
        }

+       goDestPath := os.Getenv("BUILDKIT_GO_DEST")
+       if goDestPath != "" {
+               goFiles, err := gogen.GeneratePackage(toolDescription, ppkg)
+               if err != nil {
+                       return fmt.Errorf("generating golang package: %v", err)
+               }
+
+               if err := writeFiles(filepath.Join(goDestPath, "v"+version), goFiles); err != nil {
+                       return err
+               }
+       }
+
        return nil
 }

```
