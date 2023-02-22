!(function o(r, a, i) {
    function s(t, e) {
        if (!a[t]) {
            if (!r[t]) {
                var n = "function" == typeof require && require;
                if (!e && n) return n(t, !0);
                if (l) return l(t, !0);
                throw (
                    (((e = new Error("Cannot find module '" + t + "'")).code =
                        "MODULE_NOT_FOUND"),
                        e)
                );
            }
            (n = a[t] = { exports: {} }),
                r[t][0].call(
                    n.exports,
                    function (e) {
                        return s(r[t][1][e] || e);
                    },
                    n,
                    n.exports,
                    o,
                    r,
                    a,
                    i
                );
        }
        return a[t].exports;
    }
    for (
        var l = "function" == typeof require && require, e = 0;
        e < i.length;
        e++
    )
        s(i[e]);
    return s;
})(
    {
        1: [
            function (e, t, n) {
                "use strict";
                var e = e("../functions/querySelector"),
                    o = (0, e.qS)(".show-instruction-button-js"),
                    r = (0, e.qS)(".instruction-block-js"),
                    a = (0, e.qS)(".instruction__show-img"),
                    o =
                        (o &&
                        o.addEventListener("click", function () {
                            r.classList.toggle("hide"), a.classList.toggle("active");
                        }),
                            (0, e.qSA)(".step-3__copy-text"));
                o &&
                o.forEach(function (o) {
                    o.addEventListener("click", function (e) {
                        var t = e.target.querySelector(".step-3__copy-img"),
                            n = e.target.querySelector(".step-3__copy-done"),
                            e =
                                ((t.style.display = "none"),
                                    (n.style.display = "block"),
                                    (o.style.animationName = "github-button"),
                                    o);
                        (e = o.parentElement
                            .querySelector(".step-3__item-text")
                            .textContent.trim()),
                            navigator.clipboard.writeText(e),
                            setTimeout(function () {
                                (t.style.display = "block"),
                                    (n.style.display = "none"),
                                    (o.style.animationName = "none");
                            }, 2e3);
                    });
                });
                var o = new Date(),
                    i = o.getMonth() + 1,
                    s = new Date(o.getFullYear(), o.getMonth(), 0).getDate();
                var l = i < 10 ? ".0" + i : i;
                (o = (function () {
                    for (var e = [], t = 1; t < s; t++)
                        t % 2 == 0 &&
                        e.push({
                            x: Math.floor(100 * Math.random()),
                            y: 10 <= t || 0 == t ? t + l : "0" + t + l,
                        });
                    return e;
                })()),
                    (i = {
                        series: [
                            {
                                name: "Views",
                                data: o.map(function (e) {
                                    return e.x;
                                }),
                            },
                        ],
                        chart: {
                            type: "area",
                            height: 340,
                            zoom: { enabled: !1 },
                            toolbar: { show: !1 },
                        },
                        colors: ["#6D96FF"],
                        fill: {
                            type: "gradient",
                            colors: ["#A0ACFF"],
                            gradient: { shadeIntensity: 1, opacityFrom: 0.7, opacityTo: 0.9 },
                        },
                        dataLabels: { enabled: !1 },
                        stroke: { curve: "smooth" },
                        labels: o.map(function (e) {
                            return e.y;
                        }),
                        xaxis: { type: "datey" },
                        yaxis: { opposite: !0 },
                        legend: { horizontalAlign: "left" },
                    }),
                    (o = (0, e.qS)(".chart-js")),
                o && new ApexCharts(o, i).render(),
                    (o = (0, e.qS)(".chart-js-rows"));
                o &&
                new ApexCharts(o, {
                    series: [
                        {
                            name: "Registaation",
                            data: [35, 41, 36, 26, 45, 48, 52, 53, 41],
                        },
                    ],
                    chart: {
                        type: "bar",
                        height: 350,
                        zoom: { enabled: !1 },
                        toolbar: { show: !1 },
                    },
                    colors: ["#13161B"],
                    plotOptions: {
                        bar: {
                            horizontal: !1,
                            columnWidth: "10%",
                            endingShape: "rounded",
                        },
                    },
                    dataLabels: { enabled: !1 },
                    stroke: { show: !0, width: 2, colors: ["transparent"] },
                    xaxis: {
                        categories: [
                            "Jan 03",
                            "Jan 06",
                            "Jan 09",
                            "Jan 12",
                            "Jan 15",
                            "Jan 18",
                            "Jan 24",
                            "Jan 27",
                            "Jan 30",
                        ],
                    },
                    yaxis: { opposite: !0 },
                    fill: { opacity: 1 },
                    legend: { horizontalAlign: "left" },
                }).render();
            },
            { "../functions/querySelector": 2 },
        ],
        2: [
            function (e, t, n) {
                "use strict";
                Object.defineProperty(n, "__esModule", { value: !0 }),
                    (n.qSA = n.qS = void 0);
                var o = document.querySelector.bind(document);
                n.qS = o;
                n.qSA = function (e) {
                    return Array.from(document.querySelectorAll("".concat(e)));
                };
            },
            {},
        ],
    },
    {},
    [1]
);
