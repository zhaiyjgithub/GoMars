GoMars API

    
##### Brief Description

- Launch a new car or re-control the existing car by car's name

##### Request URL
- ` http://42.192.92.99:8088/GoMarsService/Car/LaunchNow `
  
##### Request method
- POST 

##### Parameter

|Name|Required|Type|description|
|:----    |:---|:----- |-----   |
|Name |yes  |string |   |

##### Response

|Name|Type|Description|
|:-----  |:-----|-----  |
|code |int   | code = 0, is Ok; or fail|
|msg |string   | |
|data | interface | 
|data.Name | string | car's name 
|data.X | int | car's  x-axis
|data.Y | int | car's y-axis

##### Response sample
- Launch a new car by name
``` 
  {
      "code": 0,
      "data": {
          "Name": "yuanji",
          "X": 1,
          "Y": 4,
          "Direction": "North"
      },
      "msg": "You have re-control this car[yuanji]."
  }
```

- if you request an existing car by name:
```
{
    "code": 0,
    "data": {
        "Name": "yuanji",
        "X": 1,
        "Y": 4,
        "Direction": "North"
    },
    "msg": "You have re-control this car[yuanji]."
}
```

##### Notes

- none


-------

##### Brief Description

- Send command string to the car

##### Request URL
- ` http://42.192.92.99:8088/GoMarsService/Car/SendCommand `
  
##### Request method
- POST 

##### Parameter

|Name|Required|Type|description|
|:----    |:---|:----- |-----   |
|Name |yes  |string |   |
|Command|yes | string | F=move forward one step; B= move back on step; L = rotate left; R = Rotate right; H = send the helicopter to detect the surrounding

##### Response

|Name|Type|Description|
|:-----  |:-----|-----  |
|code |int   | code = 0, is Ok; or fail|
|msg |string   | |
|data | interface | 
|data.Name | string | car's name 
|data.X | int | car's  x-axis
|data.Y | int | car's y-axis
|data.Direction | string | Enum {"North", "East", "South", "West"}

##### Response sample
``` 
  {
      "code": 0,
      "data": {
          "Name": "yuanji",
          "X": 6,
          "Y": 15,
          "Direction": "North"
      },
      "msg": ""
  }
```

##### Note
- none

-------


##### Brief Description

- Get the current coverage of detecting

##### Request URL
- ` http://42.192.92.99:8088/GoMarsService/Car/GetCoverage `
  
##### Request method
- POST 

##### Parameter

|Name|Required|Type|description|
|:----    |:---|:----- |-----   |
|Name |yes  |string |   |

##### Response

|Name|Type|Description|
|:-----  |:-----|-----  |
|code |int   | code = 0, is Ok; or fail|
|msg |string   | |
|data | float64 | e.g. 0.0048 

##### Response sample
``` 
 {
     "code": 0,
     "data": 0.0048,
     "msg": "success"
 }
```

##### Note

- None


-------

##### Brief Description

- Get the current coverage of detecting

##### Request URL
- ` http://42.192.92.99:8088/GoMarsService/Car/GetCoverage `
  
##### Request method
- POST 

##### Parameter

|Name|Required|Type|description|
|:----    |:---|:----- |-----   |
|Name |yes  |string |   |

##### Response

|Name|Type|Description|
|:-----  |:-----|-----  |
|code |int   | code = 0, is Ok; or fail|
|msg |string   | |
|data | float64 | e.g. 0.0048 

##### Response sample
``` 
 {
     "code": 0,
     "data": 0.0048,
     "msg": "success"
 }
```

##### Note

- None

------

##### Brief Description

- Get the current position of car by name

##### Request URL
- ` http://42.192.92.99:8088/GoMarsService/Car/GetCurrentPosition `
  
##### Request method
- POST 

##### Parameter

|Name|Required|Type|description|
|:----    |:---|:----- |-----   |
|Name |yes  |string |   |

##### Response

|Name|Type|Description|
|:-----  |:-----|-----  |
|code |int   | code = 0, is Ok; or fail|
|msg |string   | |
|data | interface | 
|data.Name | string | car's name 
|data.X | int | car's  x-axis
|data.Y | int | car's y-axis
|data.Direction | string | Enum {"North", "East", "South", "West"}

##### Response sample
``` 
{
    "code": 0,
    "data": {
        "Name": "yuanji",
        "X": 6,
        "Y": 15,
        "Direction": "North"
    },
    "msg": "success"
}

##### Note
- None

```


