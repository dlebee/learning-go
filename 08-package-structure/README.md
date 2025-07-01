# Reading

The reason this example is being created is to help understand how package are structured.

in GoLang things are organized that if they are the same package if the name of the function starts with a lower case its private, if it starts with a capital letter its public and exported.

```golang
func PublicThing() {}  // Can be accessed from other packages
func privateThing() {} // Only accessible inside this package
```

You can see how utils is a package with a private and public function used by the main package.