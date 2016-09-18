
package chronobiology_test

import (
    "time"
    "testing"
    "reflect"
    "github.com/kelvins/chronobiology"
)

func TestInvalidParametersHigherActivity(t *testing.T) {
    // Get UTC
    utc, _ := time.LoadLocation("UTC")

    // Create the slices
    var myDateTime []time.Time
    var myData []float64

    // Call the function with empty slices
    _, _, err := chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: Empty")
    }

    tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

    // Fill the myDateTime with 1 - 12 hours
    for index := 0; index < 8; index++ {
        tempDateTime = tempDateTime.Add(1 * time.Hour)
        myDateTime = append(myDateTime, tempDateTime)
    }

    // Call the function with myData empty
    _, _, err = chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: Empty")
    }

    myData = append(myData, 450.0) // 01
    myData = append(myData, 050.0) // 02
    myData = append(myData, 025.0) // 03
    myData = append(myData, 020.0) // 04
    myData = append(myData, 100.0) // 05
    myData = append(myData, 500.0) // 06
    myData = append(myData, 250.0) // 07

    _, _, err = chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: DifferentSize")
    }

    myData = append(myData, 050.0) // 08

    _, _, err = chronobiology.HigherActivity(0, myDateTime, myData)

    if err == nil {
        t.Error("Expect: InvalidHours")
    }

    _, _, err = chronobiology.HigherActivity(20, myDateTime, myData)

    if err == nil {
        t.Error("Expect: HoursHigher")
    }
}

func TestHigherActivity(t *testing.T) {
    // Table tests
    var tTests = []struct {
        hours int
        higherActivity float64
        onsetHigherActivity string
    }{
        { 01, 990.0000, "01/01/2015 08:00:00" },
        { 02, 620.0000, "01/01/2015 07:00:00" },
        { 05, 482.0000, "01/01/2015 06:00:00" },
        { 06, 418.3333, "01/01/2015 05:00:00" },
        { 07, 364.2857, "01/01/2015 05:00:00" },
        { 10, 305.5000, "01/01/2015 01:00:00" },
    }

    // Get UTC
    utc, _ := time.LoadLocation("UTC")

    // Create the slices
    var myDateTime []time.Time
    var myData []float64

    tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

    // Fill the myDateTime with 1 - 12 hours
    for index := 0; index < 12; index++ {
        tempDateTime = tempDateTime.Add(1 * time.Hour)
        myDateTime = append(myDateTime, tempDateTime)
    }

    // Creates the data slice
    myData = append(myData, 450.0) // 01
    myData = append(myData, 050.0) // 02
    myData = append(myData, 025.0) // 03
    myData = append(myData, 020.0) // 04
    myData = append(myData, 100.0) // 05
    myData = append(myData, 500.0) // 06
    myData = append(myData, 250.0) // 07
    myData = append(myData, 990.0) // 08
    myData = append(myData, 130.0) // 09
    myData = append(myData, 540.0) // 10
    myData = append(myData, 040.0) // 11
    myData = append(myData, 050.0) // 12

    // Test with all values in the table
    for _, pair := range tTests {
        higherActivity, onsetHigherActivity, err := chronobiology.HigherActivity(pair.hours, myDateTime, myData)
        if err != nil {
            t.Error(
                "For: ", pair.hours, " hours - ",
                "expect: error not nil",
            )
        }
        if higherActivity != pair.higherActivity {
            t.Error(
                "For: ", pair.hours, " hours - ",
                "expected: ", pair.higherActivity,
                "received: ", higherActivity,
            )
        }
        if onsetHigherActivity.Format("02/01/2006 15:04:05") != pair.onsetHigherActivity {
            t.Error(
                "For: ", pair.hours, " hours - ",
                "expected: ", pair.onsetHigherActivity,
                "received: ", onsetHigherActivity.Format("02/01/2006 15:04:05"),
            )
        }
    }
}

func TestInvalidParametersLowerActivity(t *testing.T) {
    // Get UTC
    utc, _ := time.LoadLocation("UTC")

    // Create the slices
    var myDateTime []time.Time
    var myData []float64

    // Call the function with empty slices
    _, _, err := chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: Empty")
    }

    tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

    // Fill the myDateTime with 1 - 12 hours
    for index := 0; index < 8; index++ {
        tempDateTime = tempDateTime.Add(1 * time.Hour)
        myDateTime = append(myDateTime, tempDateTime)
    }

    // Call the function with myData empty
    _, _, err = chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: Empty")
    }

    myData = append(myData, 450.0) // 01
    myData = append(myData, 050.0) // 02
    myData = append(myData, 025.0) // 03
    myData = append(myData, 020.0) // 04
    myData = append(myData, 100.0) // 05
    myData = append(myData, 500.0) // 06
    myData = append(myData, 250.0) // 07

    _, _, err = chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: DifferentSize")
    }

    myData = append(myData, 050.0) // 08

    _, _, err = chronobiology.HigherActivity(0, myDateTime, myData)

    if err == nil {
        t.Error("Expect: InvalidHours")
    }

    _, _, err = chronobiology.HigherActivity(20, myDateTime, myData)

    if err == nil {
        t.Error("Expect: HoursHigher")
    }
}

func TestLowerActivity(t *testing.T) {
    // Table tests
    var tTests = []struct {
        hours int
        lowerActivity float64
        onsetLowerActivity string
    }{
        { 01, 031.5, "01/01/2016 11:00:00" },
        { 02, 061.5, "01/01/2016 10:00:00" },
        { 04, 121.5, "01/01/2016 08:00:00" },
        { 06, 181.5, "01/01/2016 06:00:00" },
        { 07, 211.5, "01/01/2016 05:00:00" },
        { 10, 301.5, "01/01/2016 02:00:00" },
    }

    // Get UTC
    utc, _ := time.LoadLocation("UTC")

    // Create the slices
    var myDateTime []time.Time
    var myData []float64

    tempDateTime := time.Date(2016,1,1,0,0,0,0,utc)

    // Fill the myDateTime (12 hours * 60 minutes) time.Minute
    for index := 0; index < (12*60); index++ {
        tempDateTime = tempDateTime.Add(1 * time.Minute)
        myDateTime = append(myDateTime, tempDateTime)
        myData = append(myData, float64((12*60)-index))
    }

    // Test with all values in the table
    for _, pair := range tTests {
        lowerActivity, onsetLowerActivity, err := chronobiology.LowerActivity(pair.hours, myDateTime, myData)
        if err != nil {
            t.Error(
                "For: ", pair.hours, " hours - ",
                "expect: error not nil",
            )
        }
        if lowerActivity != pair.lowerActivity {
            t.Error(
                "For: ", pair.hours, " hours - ",
                "expected: ", pair.lowerActivity,
                "received: ", lowerActivity,
            )
        }
        if onsetLowerActivity.Format("02/01/2006 15:04:05") != pair.onsetLowerActivity {
            t.Error(
                "For: ", pair.hours, " hours - ",
                "expected: ", pair.onsetLowerActivity,
                "received: ", onsetLowerActivity.Format("02/01/2006 15:04:05"),
            )
        }
    }
}

func TestRelativeAmplitude(t *testing.T) {
    // Expect an error
    _, err := chronobiology.RelativeAmplitude(0.0, 0.0)
    if err == nil {
        t.Error(
            "Error: ", err,
        )
    }

    // Table tests
    var tTests = []struct {
        highestAverage float64
        lowestAverage float64
        relativeAmplitude float64
    }{
        { 180.0, 050.0, 0.5652 },
        { 550.0, 125.0, 0.6296 },
        { 101.0, 100.5, 0.0025 },
        { 898.0, 315.0, 0.4806 },
        { 211.5, 075.5, 0.4739 },
        { 620.0, 020.0, 0.9375 },
        { 780.0, 010.0, 0.9747 },
    }

    // Test with all values in the table
    for _, pair := range tTests {
        relativeAmplitude, err := chronobiology.RelativeAmplitude(pair.highestAverage, pair.lowestAverage)
        if err != nil {
            t.Error(
                "Error: ", err,
            )
        }
        if relativeAmplitude != pair.relativeAmplitude {
            t.Error(
                "Expected: ", pair.relativeAmplitude,
                "Received: ", relativeAmplitude,
            )
        }
    }
}

func TestFindEpoch(t *testing.T) {

    utc, _ := time.LoadLocation("UTC")
    tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

    var dateTimeEmpty []time.Time

    var dateTime60sec []time.Time
    for index := 0; index < 420; index++ {
        tempDateTime  = tempDateTime.Add(1 * time.Minute)
        dateTime60sec = append(dateTime60sec, tempDateTime)
    }

    var dateTime30sec []time.Time
    for index := 0; index < 120; index++ {
        tempDateTime  = tempDateTime.Add(30 * time.Second)
        dateTime30sec = append(dateTime30sec, tempDateTime)
    }

    var dateTime5sec []time.Time
    for index := 0; index < 10; index++ {
        tempDateTime = tempDateTime.Add(5 * time.Second)
        dateTime5sec = append(dateTime5sec, tempDateTime)
    }

    var dateTime180sec []time.Time
    for index := 0; index < 820; index++ {
        tempDateTime   = tempDateTime.Add(3 * time.Minute)
        dateTime180sec = append(dateTime180sec, tempDateTime)
    }

    var dateTime120sec []time.Time
    for index := 0; index < 100; index++ {
        tempDateTime   = tempDateTime.Add(2 * time.Minute)
        dateTime120sec = append(dateTime120sec, tempDateTime)
    }
    for index := 0; index < 30; index++ {
        tempDateTime   = tempDateTime.Add(1 * time.Minute)
        dateTime120sec = append(dateTime120sec, tempDateTime)
    }
    for index := 0; index < 15; index++ {
        tempDateTime   = tempDateTime.Add(30 * time.Second)
        dateTime120sec = append(dateTime120sec, tempDateTime)
    }
    for index := 0; index < 100; index++ {
        tempDateTime   = tempDateTime.Add(2 * time.Minute)
        dateTime120sec = append(dateTime120sec, tempDateTime)
    }

    var dateTime360sec []time.Time
    for index := 0; index < 50; index++ {
        tempDateTime   = tempDateTime.Add(6 * time.Minute)
        dateTime360sec = append(dateTime360sec, tempDateTime)
    }
    for index := 0; index < 49; index++ {
        tempDateTime   = tempDateTime.Add(4 * time.Minute)
        dateTime360sec = append(dateTime360sec, tempDateTime)
    }
    for index := 0; index < 49; index++ {
        tempDateTime   = tempDateTime.Add(2 * time.Minute)
        dateTime360sec = append(dateTime360sec, tempDateTime)
    }

    var dateTimeInvalid []time.Time
    for index := 0; index < 250; index++ {
        dateTimeInvalid = append(dateTimeInvalid, tempDateTime)
    }

    // Table tests
    var tTests = []struct {
        dateTime []time.Time
        epoch int
    }{
        { dateTimeEmpty,    0 },
        { dateTime60sec,   60 },
        { dateTime30sec,   30 },
        { dateTime5sec,     5 },
        { dateTime180sec, 180 },
        { dateTime120sec, 120 },
        { dateTime360sec, 360 },
        { dateTimeInvalid,  0 },
    }

    // Test with all values in the table
    for _, pair := range tTests {
        epoch := chronobiology.FindEpoch(pair.dateTime)
        if epoch != pair.epoch {
            t.Error(
                "Expected: ", pair.epoch,
                "Received: ", epoch,
            )
        }
    }
}

func TestConvertDataBasedOnEpoch(t *testing.T) {

    utc, _ := time.LoadLocation("UTC")
    tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

    var dateTimeEmpty []time.Time
    var dataEmpty []float64

    _, _, err := chronobiology.ConvertDataBasedOnEpoch(dateTimeEmpty, dataEmpty, 120)

    if err == nil {
        t.Error("Expect error Empty")
    }

    var dateTimeInvalid []time.Time
    var dataInvalid []float64

    tempDateTime    = tempDateTime.Add(1 * time.Minute)
    dateTimeInvalid = append(dateTimeInvalid, tempDateTime)
    tempDateTime    = tempDateTime.Add(1 * time.Minute)
    dateTimeInvalid = append(dateTimeInvalid, tempDateTime)
    dataInvalid     = append(dataInvalid, 123.5)

    _, _, err = chronobiology.ConvertDataBasedOnEpoch(dateTimeInvalid, dataInvalid, 120)

    if err == nil {
        t.Error("Expect error DifferentSize")
    }

    var dateTime60secs []time.Time
    var data60secs []float64

    tempDateTime = time.Date(2015,1,1,0,0,0,0,utc)
    for index := 0; index < 40; index++ {
        dateTime60secs = append(dateTime60secs, tempDateTime)
        data60secs     = append(data60secs, 250.0)
        tempDateTime   = tempDateTime.Add(60 * time.Second)
    }

    var newDateTime30secs []time.Time
    var newData30secs []float64

    tempDateTime = dateTime60secs[0]
    for index := 0; index < 80; index++ {
        newDateTime30secs = append(newDateTime30secs, tempDateTime)
        newData30secs     = append(newData30secs, 250.0)
        tempDateTime      = tempDateTime.Add(30 * time.Second)
    }

    var newDateTime120secs []time.Time
    var newData120secs []float64

    tempDateTime = dateTime60secs[0]
    for index := 0; index < 20; index++ {
        newDateTime120secs = append(newDateTime120secs, tempDateTime)
        newData120secs     = append(newData120secs, 250.0)
        tempDateTime       = tempDateTime.Add(120 * time.Second)
    }

    var newDateTime90secs []time.Time
    var newData90secs []float64

    tempDateTime = dateTime60secs[0]
    for index := 0; index < 20; index++ {
        newDateTime90secs = append(newDateTime90secs, tempDateTime)
        newData90secs     = append(newData90secs, 333.3333)
        tempDateTime       = tempDateTime.Add(90 * time.Second)
    }

    var newDateTime15secs []time.Time
    var newData15secs []float64

    tempDateTime = dateTime60secs[0]
    for index := 0; index < 160; index++ {
        newDateTime15secs = append(newDateTime15secs, tempDateTime)
        newData15secs     = append(newData15secs, 250.0)
        tempDateTime      = tempDateTime.Add(15 * time.Second)
    }

    var newDateTime240secs []time.Time
    var newData240secs []float64

    tempDateTime = dateTime60secs[0]
    for index := 0; index < 10; index++ {
        newDateTime240secs = append(newDateTime240secs, tempDateTime)
        newData240secs     = append(newData240secs, 250.0)
        tempDateTime       = tempDateTime.Add(240 * time.Second)
    }

    // Table tests
    var tTests = []struct {
        dateTime []time.Time
        data []float64
        newEpoch int
        newDateTime []time.Time
        newData []float64
    }{
        { dateTime60secs, data60secs,  30,  newDateTime30secs,  newData30secs },
        { dateTime60secs, data60secs, 120, newDateTime120secs, newData120secs },
        { dateTime60secs, data60secs,  90,  newDateTime90secs,  newData90secs },
        { dateTime60secs, data60secs,  15,  newDateTime15secs,  newData15secs },
        { dateTime60secs, data60secs, 240, newDateTime240secs, newData240secs },
    }

    // Test with all values in the table
    for _, table := range tTests {
        newDateTime, newData, err := chronobiology.ConvertDataBasedOnEpoch(table.dateTime, table.data, table.newEpoch)

        if err != nil {
            t.Error("Expected error = nil.")
        }
        if !reflect.DeepEqual(newDateTime, table.newDateTime) {
            t.Error("Different dateTime slices. NewEpoch : ", table.newEpoch,)
        }
        if !reflect.DeepEqual(newData, table.newData) {
            t.Error("Different data slices. NewEpoch : ", table.newEpoch,)
        }
    }
}

func TestIntradailyVariability(t *testing.T) {

    /* TEST WITH INVALID PARAMETERS */

    utc, _ := time.LoadLocation("UTC")
    tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

    var dateTimeEmpty []time.Time
    var dataEmpty []float64

    _, err := chronobiology.IntradailyVariability(dateTimeEmpty, dataEmpty)

    if err == nil {
        t.Error("Expected error : Empty")
    }

    var dateTimeDifferentSize []time.Time
    var dataDifferentSize []float64

    dateTimeDifferentSize = append(dateTimeDifferentSize, tempDateTime)
    tempDateTime          = tempDateTime.Add(60 * time.Second)
    dateTimeDifferentSize = append(dateTimeDifferentSize, tempDateTime)
    tempDateTime          = tempDateTime.Add(60 * time.Second)

    dataDifferentSize = append(dataDifferentSize, 250.0)
    dataDifferentSize = append(dataDifferentSize, 250.0)
    dataDifferentSize = append(dataDifferentSize, 250.0)

    _, err = chronobiology.IntradailyVariability(dateTimeEmpty, dataEmpty)

    if err == nil {
        t.Error("Expected error : DifferentSize")
    }

    var dateTimeLess2Hours []time.Time
    var dataLess2Hours []float64

    dateTimeLess2Hours = append(dateTimeLess2Hours, tempDateTime)
    tempDateTime       = tempDateTime.Add(60 * time.Second)
    dateTimeLess2Hours = append(dateTimeLess2Hours, tempDateTime)
    tempDateTime       = tempDateTime.Add(60 * time.Second)
    dateTimeLess2Hours = append(dateTimeLess2Hours, tempDateTime)
    tempDateTime       = tempDateTime.Add(60 * time.Second)

    dataLess2Hours = append(dataLess2Hours, 250.0)
    dataLess2Hours = append(dataLess2Hours, 250.0)
    dataLess2Hours = append(dataLess2Hours, 250.0)

    _, err = chronobiology.IntradailyVariability(dateTimeEmpty, dataEmpty)

    if err == nil {
        t.Error("Expected error : LessThan2Hours")
    }
}
