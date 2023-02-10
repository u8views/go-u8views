import {formatDay} from "./time";

fetch("/api/v1/users/stats.json")
    .then(function (response) {
        return response.json();
    })
    .then(render)
    .catch(console.error);

function render(stats) {
    const options = {
        series: [{
            name: "Users",
            data: stats.map((item) => {
                return item.count;
            })
        }],
        chart: {
            type: 'area',
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
            }
        },
        dataLabels: {
            enabled: false
        },
        stroke: {
            curve: 'smooth'
        },
        labels: stats.map((item) => {
            return formatDay(item.time);
        }),
        xaxis: {
            type: 'datey',
        },
        yaxis: {
            opposite: true
        },
        legend: {
            horizontalAlign: 'left'
        }
    };

    const chartApex = new ApexCharts(document.querySelector(".chart-js"), options);
    chartApex.render();
}