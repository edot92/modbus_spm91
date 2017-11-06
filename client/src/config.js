/* no es-lint */
var baseUrl;
if (process.env.NODE_ENV === "development") {
  baseUrl = "http://localhost:9000/api";
  // headers.common['Access-Control-Allow-Origin'] =
  //   'http://127.0.0.1:9000';
  // headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
} else {
  baseUrl = "/api";
}

const config = {
  locale: "en-US", // en-US, zh-CN
  url: baseUrl,
  ajaxUploadUrl: baseUrl + "/admin/api/upload",
  debug: {
    mock: false, // enable mock
    http: false // http request log
  },
  api: baseUrl
  // locale: 'en', //en
  // api: 'http://192.168.1.108:3333/admin/api'
};

global.config = config;

export default config;
