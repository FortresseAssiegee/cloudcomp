export function getNowTime() {
    let d = new Date();
    let nowTime =
        d.getFullYear() +
        "-" +
        (d.getMonth() + 1) +
        "-" +
        d.getDate() +
        " " +
        addZero(d.getHours()) +
        ":" +
        addZero(d.getMinutes()) +
        ":" +
        addZero(d.getSeconds());
    return nowTime;

}
export function getNowDate() {
    let d = new Date();
    let nowTime =
        d.getFullYear() +
        "-" +
        (d.getMonth() + 1) +
        "-" +
        d.getDate()
    return nowTime;

}

// 计算现在的时间过某个时间后
export function GetDuring(during) {
    let d = new Date();


    let nowTime = d.getFullYear() +
        "-" +
        (d.getMonth() + 1) +
        "-" +
        d.getDate() +
        " " + addZero(d.getHours()) +
        ":" +
        addZero(d.getMinutes()) +
        ":" +
        addZero(d.getSeconds());


    var x = new Date(d.setMinutes(d.getMinutes() + during * 1));
    let endTime = x.getFullYear() +
        "-" +
        (x.getMonth() + 1) +
        "-" +
        x.getDate() +
        " " + addZero(x.getHours()) +
        ":" +
        addZero(x.getMinutes()) +
        ":" +
        addZero(x.getSeconds());

    return {
        now: nowTime,
        end: endTime
    }
}

function addZero(s) {
    return s < 10 ? "0" + s : s;
}