import {formatDay} from "./time";
import {initInstruction} from "./instruction";
import {initCopyCodeButtons} from "./copy-button";

const socialProviderUserId = document.body.getAttribute("data-current-page-profile-social-provider-user-id");

fetch(`/api/v1/github/profiles/${socialProviderUserId}/views/stats.json`)
    .then(function (response) {
        return response.json();
    })
    .then(renderViewsStatistics)
    .catch(console.error);

fetch(`/api/v1/github/profiles/${socialProviderUserId}/referrals/stats.json`)
    .then(function (response) {
        return response.json();
    })
    .then(renderReferralRegistrationStatistics)
    .catch(console.error);

function renderViewsStatistics(stats) {
    stats = groupByDay(stats);

    const options = {
        series: [{
            name: "Views",
            data: stats.map((item) => {
                return item.count;
            })
        }],
        chart: {
            type: "area",
            height: 340,
            zoom: {
                enabled: false
            },
            toolbar: {
                show: false
            }
        },
        colors: ["black"],
        fill: {
            type: "gradient",
            colors: ['#A0ACFF'],
            gradient: {
                shadeIntensity: 1,
                opacityFrom: 0.7,
                opacityTo: 0.9,
            },
        },
        dataLabels: {
            enabled: false
        },
        stroke: {
            curve: "smooth"
        },
        labels: stats.map((item) => {
            return item.time;
        }),
        xaxis: {
            type: "datey",
        },
        yaxis: {
            opposite: true
        },
        legend: {
            horizontalAlign: "left"
        }
    };

    const chartApex = new ApexCharts(document.querySelector(".js-chart-views-statistic"), options);
    chartApex.render();
}

function renderReferralRegistrationStatistics(stats) {
    const options = {
        series: [{
            name: "Users",
            data: stats.map((item) => {
                return item.count;
            })
        }],
        chart: {
            type: "bar",
            height: 350,
            zoom: {
                enabled: false
            },
            toolbar: {
                show: false
            }
        },
        colors: ["#13161B"],
        plotOptions: {
            bar: {
                horizontal: false,
                columnWidth: "10%",
                endingShape: "rounded"
            },
        },
        dataLabels: {
            enabled: false
        },
        stroke: {
            show: true,
            width: 2,
            colors: ["transparent"]
        },
        xaxis: {
            categories: stats.map((item) => {
                return formatDay(item.time);
            }),
        },
        yaxis: {
            opposite: true
        },
        fill: {
            opacity: 1
        },
        legend: {
            horizontalAlign: "left"
        }
    };

    const chartApex = new ApexCharts(document.querySelector(".js-chart-referral-registration-statistics"), options);
    chartApex.render();
}

function groupByDay(rows) {
    const result = {};

    for (const row of rows) {
        const time = row.time - (row.time % 86400);

        if (result.hasOwnProperty(time)) {
            result[time].count += row.count;
        } else {
            result[time] = {
                time: formatDay(time),
                count: row.count,
            };
        }
    }

    return Object.values(result);
}

initInstruction();
initCopyCodeButtons();