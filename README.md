# Parking Lot

Shell app to simulate parking system

![image](https://user-images.githubusercontent.com/5858756/56428038-330dfc00-62e8-11e9-8a79-8b3261062f74.png)

## how to play

run install/download depedencies  & unit testing

```bash
bin/setup
```

run functional spec

```bash
bin/run_functional_tests
```

run shell app

```bash
bin/parking_lot
```

### project structure

```
src/
    handlers/
        parking.go
    models/
        car.go
        parking.go
    service/
        parking.go
    main.go
go.mod
Makefile
```

## checkpoint
- [X] list feature of parking lot app
    - [x] create parking slot
    - [x] cek status parking slot
    - [X] new car come to park
    - [x] car exit / leave
    - [x] search car by colour
    - [x] search slot by car colour
    - [x] search slot by car platno
    - [X] input from typing
    - [X] input from file
    - [x] exit from app
- [X] git commit
- [X] readme
- [ ] unit testing
    - [X] helpers
    - [ ] handlers
    - [ ] services


### functional test result
```bash
example_id                       | status  | run_time        |
-------------------------------- | ------- | --------------- |
./spec/end_to_end_spec.rb[1:1:1] | passed  | 0.05352 seconds |
./spec/end_to_end_spec.rb[1:1:2] | passed  | 12.12 seconds   |
./spec/parking_lot_spec.rb[1:1]  | passed  | 0.8121 seconds  |
./spec/parking_lot_spec.rb[1:2]  | passed  | 1.12 seconds    |
./spec/parking_lot_spec.rb[1:3]  | passed  | 1.42 seconds    |
./spec/parking_lot_spec.rb[1:4]  | passed  | 2.03 seconds    |
./spec/parking_lot_spec.rb[1:5]  | pending | 0.00004 seconds |
```
