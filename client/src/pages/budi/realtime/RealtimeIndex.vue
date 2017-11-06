<template>
  <div>
    <div v-if="error!=''">
      <v-card>
        Error
        <v-card-text>
          {{error}}
        </v-card-text>
      </v-card>
    </div>
      <v-card>

    <v-layout row wrap>
      <v-flex md6 xs12>
    <div id="gaugeTegangan"></div>
      </v-flex>
        <v-flex md6 xs12>
    <div id="gaugeArus"></div>
      </v-flex>
    </v-layout>
      </v-card>
      <v-card>
    <v-layout>
      <v-flex md6 xs12>
        <div id="grafikKWH" style="min-width: 200px; height: 400px; margin: 0 auto"></div>
      </v-flex>
        <v-flex md6 xs12>
        <div id="grafikKW" style="min-width: 200px; height: 400px; margin: 0 auto"></div>
      </v-flex>
    </v-layout>
      </v-card>

  </div>
</template>
<script>
export default {
  data () {
    return {
      error: '',
      gaugeTegangan: {},
      gaugeArus: {},
      nilaiKWH: 0.0,
      nilaiKW: 0.0
    };
  },
  mounted () {
    const thisVue = this;
    $(document).ready(function () {
      thisVue.ChartInitKWH();
      thisVue.ChartInitKW();
      thisVue.gaugeInit();
      setInterval(function () {
        thisVue.getdataSPM();
      }, 1000);
    });
  },
  methods: {
    // socketIO () {
    //   var socket = window.io('http://localhost');
    //   socket.on('connect', function () {
    //     console.log('connect');
    //   });
    //   socket.on('event', function (data) {
    //     console.log(data);
    //   });
    //   socket.on('disconnect', function () {
    //     console.log('diskonek');
    //   });
    // },
    gaugeInit () {
      const thisVue = this;
      // init tegangan,arus , daya
      thisVue.gaugeInitTegangan();
      thisVue.gaugeInitArus();
    },
    getdataSPM () {
      var thisVue = this;
      // console.log(thisVue.gaugeTegangan);
      thisVue
        .$http({
          method: 'get',
          url: '/energymeter'
        })
        .then(res => {
          // console.log(res.data.payload.Tegangan);
          if (res.data.error === false) {
            thisVue.error = '';
            thisVue.gaugeTegangan.update(parseFloat(res.data.payload.Tegangan));
            thisVue.gaugeArus.update(parseFloat(res.data.payload.Arus));
            thisVue.nilaiKWH = parseFloat(res.data.payload.DayaAktif);
            thisVue.nilaiKW = parseFloat(res.data.payload.Daya);
          } else {
            thisVue.error = res.data.message;
          }
        })
        .catch(err => {
          console.error(err);
        });
    },
    gaugeInitTegangan () {
      const thisVue = this;
      window.Highcharts.chart(
        'gaugeTegangan',
        {
          chart: {
            type: 'gauge',
            plotBackgroundColor: null,
            plotBackgroundImage: null,
            plotBorderWidth: 0,
            plotShadow: false
          },

          title: {
            text: 'Tegangan'
          },
          pane: {
            startAngle: -150,
            endAngle: 150,
            background: [
              {
                backgroundColor: {
                  linearGradient: { x1: 0, y1: 0, x2: 0, y2: 1 },
                  stops: [[0, '#FFF'], [1, '#333']]
                },
                borderWidth: 0,
                outerRadius: '109%'
              },
              {
                backgroundColor: {
                  linearGradient: { x1: 0, y1: 0, x2: 0, y2: 1 },
                  stops: [[0, '#333'], [1, '#FFF']]
                },
                borderWidth: 1,
                outerRadius: '107%'
              },
              {
                // default background
              },
              {
                backgroundColor: '#DDD',
                borderWidth: 0,
                outerRadius: '105%',
                innerRadius: '103%'
              }
            ]
          },
          // the value axis
          yAxis: {
            min: 0.0,
            max: 300.0,
            minorTickInterval: 'auto',
            minorTickWidth: 1,
            minorTickLength: 10,
            minorTickPosition: 'inside',
            minorTickColor: '#666',

            tickPixelInterval: 30,
            tickWidth: 2,
            tickPosition: 'inside',
            tickLength: 10,
            tickColor: '#666',
            labels: {
              step: 2,
              rotation: 'auto'
            },
            title: {
              text: 'Volt AC'
            },
            plotBands: [
              {
                from: 0,
                to: 120,
                color: '#55BF3B' // green
              },
              {
                from: 210,
                to: 239,
                color: '#DDDF0D' // yellow
              },
              {
                from: 240,
                to: 300,
                color: '#DF5353' // red
              }
            ]
          },

          series: [
            {
              name: 'Tegangan',
              data: [0],
              tooltip: {
                valueSuffix: ' Volt AC'
              }
            }
          ]
        },
        // register variabel chart ke vue
        function (chart) {
          if (!chart.renderer.forExport) {
            thisVue.gaugeTegangan = chart.series[0].points[0];
          }
        }
      );
    },
    gaugeInitArus () {
      const thisVue = this;
      window.Highcharts.chart(
        'gaugeArus',
        {
          chart: {
            type: 'gauge',
            plotBackgroundColor: null,
            plotBackgroundImage: null,
            plotBorderWidth: 0,
            plotShadow: false
          },

          title: {
            text: 'Arus'
          },
          pane: {
            startAngle: -150,
            endAngle: 150,
            background: [
              {
                backgroundColor: {
                  linearGradient: { x1: 0, y1: 0, x2: 0, y2: 1 },
                  stops: [[0, '#FFF'], [1, '#333']]
                },
                borderWidth: 0,
                outerRadius: '109%'
              },
              {
                backgroundColor: {
                  linearGradient: { x1: 0, y1: 0, x2: 0, y2: 1 },
                  stops: [[0, '#333'], [1, '#FFF']]
                },
                borderWidth: 1,
                outerRadius: '107%'
              },
              {
                // default background
              },
              {
                backgroundColor: '#DDD',
                borderWidth: 0,
                outerRadius: '105%',
                innerRadius: '103%'
              }
            ]
          },
          // the value axis
          yAxis: {
            min: 0.0,
            max: 5.0,

            minorTickInterval: 'auto',
            minorTickWidth: 1,
            minorTickLength: 10,
            minorTickPosition: 'inside',
            minorTickColor: '#666',

            tickPixelInterval: 30,
            tickWidth: 2,
            tickPosition: 'inside',
            tickLength: 10,
            tickColor: '#666',
            labels: {
              step: 2,
              rotation: 'auto'
            },
            title: {
              text: 'Ampere'
            },
            plotBands: [
              {
                from: 0.0,
                to: 2.9,
                color: '#55BF3B' // green
              },
              {
                from: 3,
                to: 4.0,
                color: '#DDDF0D' // yellow
              },
              {
                from: 4.5,
                to: 5,
                color: '#DF5353' // red
              }
            ]
          },

          series: [
            {
              name: 'Arus',
              data: [0],
              tooltip: {
                valueSuffix: 'Ampere'
              }
            }
          ]
        },
        // register variabel chart ke vue
        function (chart) {
          if (!chart.renderer.forExport) {
            thisVue.gaugeArus = chart.series[0].points[0];
          }
        }
      );
    },
    ChartInitKWH () {
      const thisVue = this;
      window.Highcharts.setOptions({
        global: {
          useUTC: false
        }
      });

      window.Highcharts.chart('grafikKWH', {
        chart: {
          type: 'spline',
          animation: window.Highcharts.svg, // don't animate in old IE
          marginRight: 10,
          events: {
            load: function () {
              // set up the updating of the chart each second
              var series = this.series[0];
              setInterval(function () {
                var x = new Date().getTime(); // current time
                var y = thisVue.nilaiKWH;
                series.addPoint([x, y], true, true);
              }, 1000);
            }
          }
        },
        title: {
          text: 'Daya Terpakai (KWH)'
        },
        xAxis: {
          type: 'datetime',
          tickPixelInterval: 150
        },
        yAxis: {
          title: {
            text: 'KWH'
          },
          plotLines: [
            {
              value: 0,
              width: 1,
              color: '#808080'
            }
          ]
        },
        tooltip: {
          formatter: function () {
            return (
              '<b>' +
              this.series.name +
              '</b><br/>' +
              window.Highcharts.dateFormat('%Y-%m-%d %H:%M:%S', this.x) +
              '<br/>' +
              window.Highcharts.numberFormat(this.y, 2)
            );
          }
        },
        legend: {
          enabled: true
        },
        exporting: {
          enabled: true
        },
        series: [
          {
            name: 'Nilai KWH (PEMAKAIAN DAYA LISTRIK)',
            data: (function () {
              // generate an array of random data
              var data = [];
              var time = new Date().getTime();
              var i;

              for (i = -19; i <= 0; i += 1) {
                data.push({
                  x: time + i * 1000,
                  y: 0
                });
              }
              return data;
            })()
          }
        ]
      });
    },
    ChartInitKW () {
      const thisVue = this;
      window.Highcharts.setOptions({
        global: {
          useUTC: false
        }
      });

      window.Highcharts.chart('grafikKW', {
        chart: {
          type: 'spline',
          animation: window.Highcharts.svg, // don't animate in old IE
          marginRight: 10,
          events: {
            load: function () {
              // set up the updating of the chart each second
              var series = this.series[0];
              setInterval(function () {
                var x = new Date().getTime(); // current time
                var y = thisVue.nilaiKW;
                series.addPoint([x, y], true, true);
              }, 1000);
            }
          }
        },
        title: {
          text: 'Daya Aktif'
        },
        xAxis: {
          type: 'datetime',
          tickPixelInterval: 150
        },
        yAxis: {
          title: {
            text: 'KW'
          },
          plotLines: [
            {
              value: 0,
              width: 1,
              color: '#808080'
            }
          ]
        },
        tooltip: {
          formatter: function () {
            return (
              '<b>' +
              this.series.name +
              '</b><br/>' +
              window.Highcharts.dateFormat('%Y-%m-%d %H:%M:%S', this.x) +
              '<br/>' +
              window.Highcharts.numberFormat(this.y, 2)
            );
          }
        },
        legend: {
          enabled: true
        },
        exporting: {
          enabled: true
        },
        series: [
          {
            name: 'Nilai KW (Daya Aktif)',
            data: (function () {
              // generate an array of random data
              var data = [];
              var time = new Date().getTime();
              var i;

              for (i = -19; i <= 0; i += 1) {
                data.push({
                  x: time + i * 1000,
                  y: 0
                });
              }
              return data;
            })()
          }
        ]
      });
    }
  }
};
</script>

