# PaintLayering
Should find what's the layer choices for you based on the base color you provide.

```
Usage of PaintLayering:
  -file string
    	a YAML config file
```

This YAML config file should be build like this :

```
description: "Order Base Paint and their layer"
gwpaint: 
  kantor blue: 
    - red wine
    - sweet
  yellow: 
    - yellow
    - green
papaint: 
  blue: 
    - red
    - sweet
  yellow: 
    - yellow
    - green
```
