<!DOCTYPE html>

<html lang="en-US">
  <head>
    <title>Thống kê số lượng request của các API trong hệ thống</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta property="og:image" content="/static/img/logo_qb.jpg" />
    <link rel="shortcut icon" href="/static/img/logo_qb.jpg" type="image/x-icon" />
    <meta property="og:title" content="Thống kê số lượng request của các API trong hệ thống">
    <link rel="stylesheet" href="/static/css/charting.css">
    <link rel="stylesheet" href="/static/css/bootstrap.min.css">
    <link rel="stylesheet" href="/static/css/bootstrap-theme.min.css">
    <script src="/static/js/jquery-3.3.1.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="/static/js/Chart.min.js"></script>
    <script src="/static/js/charting.js"></script>
    <style>
    </style>
  </head>

  <body>
    <h2 style="text-align: center;">Thống kê số lượng request của các API trong hệ thống</h2>
    <hr>
    <div class="container">
      <div class="row">
          <div class="col-lg-12 col-md-12 col-sm-12">
            <h4 style="text-align: center;">Top 30 request được gọi nhiều nhất</h4>
            <div class="chart-container-over">
              <canvas id="chartOver1" ></canvas>
            </div>
          </div>
      </div>
      <div class="row">
          <div class="col-lg-12 col-md-12 col-sm-12">
            <h4 style="text-align: center;">Top 30 action được gọi nhiều nhất</h4>
            <div class="chart-container-over">
              <canvas id="chartOver2"></canvas>   
            </div>
          </div>
      </div>
      <hr>
      <h4 style="text-align: center;">Kê khai tất cả các request, action với các option</h4>
      <div class="row">
        <div class="col-md-3 col-lg-3 col-sm-3">
            <select class="form-control" id="type_option_1" onchange="selectType1Changed(this)">
              <!-- <option value=""> -- Chọn -- </option> -->
              <option value="request" selected="selected" >Requests</option>
              <option value="action">Actions</option>
            </select>
        </div>
        <div class="col-md-3 col-lg-3 col-sm-3">
            <select class="form-control" id="gen_hour_begin" onchange="selectTypeBeginTimeChanged(this)">
                <option value="0" selected="selected">Bắt đầu từ 0 giờ</option>
                <option value="1">Bắt đầu từ 1 giờ</option>
                <option value="2">Bắt đầu từ 2 giờ</option>
                <option value="3">Bắt đầu từ 3 giờ</option>
                <option value="4">Bắt đầu từ 4 giờ</option>
                <option value="5">Bắt đầu từ 5 giờ</option>
                <option value="6">Bắt đầu từ 6 giờ</option>
                <option value="7">Bắt đầu từ 7 giờ</option>
                <option value="8">Bắt đầu từ 8 giờ</option>
                <option value="9">Bắt đầu từ 9 giờ</option>
                <option value="10">Bắt đầu từ 10 giờ</option>
                <option value="11">Bắt đầu từ 11 giờ</option>
                <option value="12">Bắt đầu từ 12 giờ</option>
                <option value="13">Bắt đầu từ 13 giờ</option>
                <option value="14">Bắt đầu từ 14 giờ</option>
                <option value="15">Bắt đầu từ 15 giờ</option>
                <option value="16">Bắt đầu từ 16 giờ</option>
                <option value="17">Bắt đầu từ 17 giờ</option>
                <option value="18">Bắt đầu từ 18 giờ</option>
                <option value="19">Bắt đầu từ 19 giờ</option>
                <option value="20">Bắt đầu từ 20 giờ</option>
                <option value="21">Bắt đầu từ 21 giờ</option>
                <option value="22">Bắt đầu từ 22 giờ</option>
                <option value="23">Bắt đầu từ 23 giờ</option>
            </select>
        </div>
        <div class="col-md-3 col-lg-3 col-sm-3">
            <select class="form-control" id="gen_hour_end" onchange="selectTypeEndTimeChanged(this)">
                <option value="0">Kết thúc đến hết 0 giờ</option>
                <option value="1">Kết thúc đến hết 1 giờ</option>
                <option value="2">Kết thúc đến hết 2 giờ</option>
                <option value="3">Kết thúc đến hết 3 giờ</option>
                <option value="4">Kết thúc đến hết 4 giờ</option>
                <option value="5">Kết thúc đến hết 5 giờ</option>
                <option value="6">Kết thúc đến hết 6 giờ</option>
                <option value="7">Kết thúc đến hết 7 giờ</option>
                <option value="8">Kết thúc đến hết 8 giờ</option>
                <option value="9">Kết thúc đến hết 9 giờ</option>
                <option value="10">Kết thúc đến hết 10 giờ</option>
                <option value="11">Kết thúc đến hết 11 giờ</option>
                <option value="12">Kết thúc đến hết 12 giờ</option>
                <option value="13">Kết thúc đến hết 13 giờ</option>
                <option value="14">Kết thúc đến hết 14 giờ</option>
                <option value="15">Kết thúc đến hết 15 giờ</option>
                <option value="16">Kết thúc đến hết 16 giờ</option>
                <option value="17">Kết thúc đến hết 17 giờ</option>
                <option value="18">Kết thúc đến hết 18 giờ</option>
                <option value="19">Kết thúc đến hết 19 giờ</option>
                <option value="20">Kết thúc đến hết 20 giờ</option>
                <option value="21">Kết thúc đến hết 21 giờ</option>
                <option value="22">Kết thúc đến hết 22 giờ</option>
                <option value="23" selected="selected">Kết thúc đến hết 23 giờ</option>
            </select>
        </div>
        <div class="col-md-3 col-lg-3 col-sm-3">
            <select class="form-control" id="type_option_2" onchange="selectType2Changed(this)">
              <!-- <option value=""> -- Chọn -- </option> -->
              <option value="count">Số lượng</option>
              <option value="time_all">Tổng thời gian</option>
              <option value="time_average" selected="selected">Thời gian trung bình</option>
            </select>
        </div>
      </div>
      <div class="row">
        <div class="col-lg-12 col-md-12 col-sm-12 chart-container">
          <canvas id="myChart"></canvas>
        </div>
      </div>
      <hr>
      <h4 style="text-align: center;">Xem chi tiết theo dòng thời gian</h4>
      <div class="row">
        <div class="col-md-3 col-lg-3 col-sm-3">
          <select class="form-control" id="type_option_3" onchange="selectType3Changed(this)">
            <!-- <option value=""> -- Chọn -- </option> -->
            <option value="request" selected="selected" >Requests</option>
            <option value="action">Actions</option>
          </select>
        </div>
        <div class="col-md-3 col-lg-3 col-sm-3">
          <select class="form-control" id="type_option_request" onchange="selectRequestChanged(this)">
            <script>
              function LoadDataRequestAndActionSelect(){
                console.log(all_action)
                var html = '';
                var  value = $("#type_option_3").val();
                console.log(value)
                var all_request = {{.all_request}};
                var all_action = {{.all_action}};
                if (value === 'request'){
                  for (var i = 0;i < all_request.length;i++){
                    html += `<option value="` + all_request[i] + `">` + all_request[i] + `</option><\br>`;
                  }
                } 
                else if (value === 'action') {
                  for (var i = 0;i < all_action.length;i++){
                    html += `<option value="` + all_action[i] + `">` + all_action[i] + `</option><\br>`;
                  }
                }
                //document.write(html);
                document.getElementById("type_option_request").innerHTML = '';
                document.getElementById("type_option_request").innerHTML = html;
              }
              LoadDataRequestAndActionSelect();
            </script>
            <!-- {{ range $request := .all_request }}
              <option value="{{ $request }}"> {{ $request }} </option>
            {{ end }} -->
          </select>
        </div>
        <div class="col-md-3 col-lg-3 col-sm-3">
          <select class="form-control" id="type_option_hour" onchange="selectTypeHourChanged(this)">
            <!-- <option value=""> -- Chọn -- </option> -->
            <option value="hour" selected="selected">Giờ</option>
            <option value="minute">Phút</option>
            <option value="time_max">Thời gian xử lí cao nhất</option>
          </select>
        </div>
      </div>
      <div class="row">
        <div class="col-lg-12 col-md-12 col-sm-12">
          <canvas id="chartRequest" style="max-width: 1000px;"></canvas>
        </div>
      </div>
    </div>

    
    <script type="text/javascript">
 
          // static 
          var myChart = null;
          var chartOver = null;
          var chartOver2 = null;
          var charRequest = null;
          var viewAction = false;
          var viewAction3 = false;
          var option3Select = 3;
          var requestSelect = '';
          var requestMinute = 1;
          var beginTimeShowBar = 0;
          var endTimeShowBar = 23;

          // data 
          var best_request_label_count= {{.best_request_label_count}};
          var best_request_count = {{.best_request_count}};
          var best_action_label_count = {{.best_action_label_count}};
          var best_action_count = {{.best_action_count}};
          var all_action = {{.all_action}};
          var all_request ={{.all_request}};
          var action_data_count = {{.action_data_count}};
          var request_data_count = {{.request_data_count}};
          var action_data_hours = {{.action_data_hours}};
          var request_data_hours = {{.request_data_hours}};
          var the_average_time_action = {{.the_average_time_action}};
          var the_average_time_request = {{.the_average_time_request}};
          var labelLoadMinute =  {{.labelLoadMinute}};
          var request_data_minutes = {{ .request_data_minutes}};
          var labelLoadHour =  {{.labelLoadHour}};
          var request_data_hours_line = {{ .request_data_hours_line}};
          var action_data_minutes = {{ .action_data_minutes}};
          var action_data_hours_line = {{ .action_data_hours_line}};

          // option ajax
          var url_data = {{.url_data}};
          var time_load = {{.time_load}};

          setInterval(function(){ 
            $.ajax({
              type: 'GET', //get method
              url: url_data, // url
              dataType: "json",
              success: function (result)
              {
                  console.log(result);
                  if(chartOver!=null){
                    chartOver.destroy();
                  }
                  if(chartOver2!=null){
                    chartOver2.destroy();
                  }

                  console.log(result.best_count_monitor.best_request_count);
                 
                  //console.log(chartOver);
                  console.log(chartOver.data.labels);
                  //console.log(chartOver.data.datasets[0].data);

                  var defaultBestItemActionOver = 30;
                  var defaultBestItemRequestOver = 30;
                  if(parseInt(defaultBestItemRequestOver) > parseInt(result.best_count_monitor.best_request_label_count.length)){
                    defaultBestItemRequestOver = parseInt(result.best_count_monitor.best_request_label_count.length)
                  }
                  if(parseInt(defaultBestItemActionOver) > parseInt(result.best_count_monitor.best_action_label_count.length)){
                    defaultBestItemActionOver = parseInt(result.best_count_monitor.best_action_label_count.length)
                  }
                  LoadCharOver(result.best_count_monitor.best_request_label_count.slice(0,defaultBestItemRequestOver),result.best_count_monitor.best_request_count.slice(0,defaultBestItemRequestOver));
                  LoadCharOver2(result.best_count_monitor.best_action_label_count.slice(0,defaultBestItemActionOver),result.best_count_monitor.best_action_count.slice(0,defaultBestItemActionOver)); 
                  console.log(chartOver.data.labels);

                  all_action = result.all_action;
                  all_request = result.all_request;
                  action_data_count = result.monitor_data.action_data_count;
                  request_data_count = result.monitor_data.request_data_count;
                  action_data_hours = result.monitor_data.action_data_hours;
                  request_data_hours = result.monitor_data.request_data_hours;
                  the_average_time_action = result.the_average_time_monitor.the_average_time_action;
                  the_average_time_request = result.the_average_time_monitor.the_average_time_request;

                  request_data_minutes = result.monitor_data.request_data_by_minutes;
                  request_data_hours_line = result.monitor_data.request_data_by_hours;
                  action_data_minutes = result.monitor_data.action_data_by_minutes;
                  action_data_hours_line = result.monitor_data.action_data_by_hours;

                  CheckLoadCharBar();
                  CheckSelectRequest();
              }
          });
          }, time_load*1000);


          // LOAD BIỂU ĐỒ TRÒN THỐNG KÊ TẤT CẢ SỐ LƯỢNG REQUEST
          var defaultBestItemActionOver = 30;
          var defaultBestItemRequestOver = 30;
          if(parseInt(defaultBestItemRequestOver) > parseInt(best_request_label_count.length)){
            defaultBestItemRequestOver = parseInt(best_request_label_count.length)
          }
          if(parseInt(defaultBestItemActionOver) > parseInt(best_action_label_count.length)){
            defaultBestItemActionOver = parseInt(best_action_label_count.length)
          }
          this.LoadCharOver(best_request_label_count.slice(0,defaultBestItemRequestOver),best_request_count.slice(0,defaultBestItemRequestOver)) 
          this.LoadCharOver2(best_action_label_count.slice(0,defaultBestItemActionOver),best_action_count.slice(0,defaultBestItemActionOver)) 

          // option xem all requests or actions
          function selectType1Changed(obj){
            var value = obj.value;
            if (value === 'request'){
              viewAction = false;
            } else if (value=== 'action'){
              viewAction = true;
            }
            this.CheckLoadCharBar();
          }
          function selectTypeBeginTimeChanged(obj){
            beginTimeShowBar = parseInt(obj.value);
            console.log(beginTimeShowBar)
            this.CheckLoadCharBar();
          }
          function selectTypeEndTimeChanged(obj){
            endTimeShowBar = parseInt(obj.value);
            console.log(endTimeShowBar)
            this.CheckLoadCharBar();
          }
          function selectType2Changed(obj){
            var value = obj.value;
            if (value === 'count'){
              option3Select = 1;
            } 
            else if (value === 'time_all')
            { 
              option3Select = 2;
            } else if (value === 'time_average')
            { 
              option3Select = 3; 
            } 
            this.CheckLoadCharBar();
          }
          function CheckLoadCharBar(){
            var max_time_action = [];
            var max_time_request = [];
            // xử lý action
            for(var actionStr = 0;actionStr < all_action.length;actionStr ++){
              var actionName  = all_action[actionStr];
              action_data_count[actionStr] = 0;
              action_data_hours[actionStr] = 0;
              var biggesttime = 0;
              for(var i = beginTimeShowBar;i<= endTimeShowBar;i++){
                action_data_count[actionStr] +=  action_data_hours_line[actionName][i].totalcount;
                action_data_hours[actionStr] +=  action_data_hours_line[actionName][i].totaltime;
                if ( biggesttime <  action_data_hours_line[actionName][i].biggesttime){
                  biggesttime = action_data_hours_line[actionName][i].biggesttime;
                }
              }
              max_time_action.push(biggesttime);
              if (action_data_count[actionStr] > 0){
                the_average_time_action[actionStr] = parseInt(action_data_hours[actionStr]/action_data_count[actionStr]);
              } else {
                the_average_time_action[actionStr] = 0;
              }
            }

            // xử lý request
            for(var requestStr = 0;requestStr < all_request.length;requestStr ++){
              var requestName  = all_request[requestStr];
              request_data_count[requestStr] = 0;
              request_data_hours[requestStr] = 0;
              var biggesttime = 0;
              for(var i = beginTimeShowBar;i<= endTimeShowBar;i++){
                request_data_count[requestStr]  +=  request_data_hours_line[requestName][i].totalcount;
                request_data_hours[requestStr] +=  request_data_hours_line[requestName][i].totaltime;
                if ( biggesttime <   request_data_hours_line[requestName][i].biggesttime){
                  biggesttime = request_data_hours_line[requestName][i].biggesttime;
                }
              }
              max_time_request.push(biggesttime);
              if (request_data_count[requestStr] > 0){
                the_average_time_request[requestStr] = parseInt(request_data_hours[requestStr]/request_data_count[requestStr]);
              } else {
                the_average_time_request[requestStr] = 0;
              }
            }
            
            // xử lý logic handle giao diện
            if(option3Select === 1){
              if (viewAction === true) {
                this.LoadChartBar(all_action,action_data_count,[],'#số lần gọi: ','lần')
              } else {
                this.LoadChartBar(all_request,request_data_count,[],'#số lần gọi: ','lần')
              }
            } else if (option3Select === 2){
              if (viewAction === true) 
              {
                this.LoadChartBar(all_action,action_data_hours,[],'#tổng thời gian thực hiện: ','nanosecond')
              } 
              else {
                this.LoadChartBar(all_request,request_data_hours,[],'#tổng thời gian thực hiện: ','nanosecond')
              }
            } else if (option3Select === 3 ){
                if (viewAction === true) 
                {
                  this.LoadChartBar(all_action,the_average_time_action,max_time_action,'#thời gian trung bình thực hiện: ','nanosecond')
                } 
                else {
                  this.LoadChartBar(all_request,the_average_time_request,max_time_request,'#thời gian trung bình thực hiện: ','nanosecond')
                }
            } 
          }
        
          // set option default
          if(requestMinute === 1) {
            if (requestSelect === '' && viewAction3 === false) {
              requestSelect = all_request[0]
            } else if (requestSelect === '' && viewAction3 === true){
              requestSelect = all_action[0]
            }
          } else if (requestMinute === 2) {
            if (requestSelect === '') {
              requestSelect =  all_request[0]
            } else if (requestSelect === '' && viewAction3 === true){
              requestSelect = all_action[0]
            }
          } else if (requestMinute === 3) {
            if (requestSelect === '') {
              requestSelect =  all_request[0]
            } else if (requestSelect === '' && viewAction3 === true){
              requestSelect = all_action[0]
            }
          }
       
          // option xem chi tiết biểu đồ
          function selectType3Changed(obj){
            var value = obj.value;
            if (value === 'request'){
              viewAction3 = false;
            } else if (value=== 'action'){
              viewAction3 = true;
            }
            LoadDataRequestAndActionSelect();
            requestSelect = '';
            this.CheckSelectRequest();
          }
          function selectRequestChanged(obj){
            var value = obj.value;
            if (value !== ''){
              requestSelect = value
              this.CheckSelectRequest();
            }
          }
          function selectTypeHourChanged(obj) {
            var value = obj.value;
            if (value === 'minute'){
              requestMinute = 2;
            } 
            else if (value === 'hour'){
              requestMinute = 1;
            }  else if (value === 'time_max'){
              requestMinute = 3;
            }
            this.CheckSelectRequest();
          }
          function CheckSelectRequest(){
            
              if (viewAction3 === false){
                if (requestSelect === ''){
                  requestSelect = all_request[0]
                }
                if(requestMinute === 2) {
                  //var keyData = requestSelect.toString();
                  var dataAll = request_data_minutes;
      
                  var dataLoad = [];
                  var labelLoadHourMin = [];
                  var i = 0;
                  for (i = 0; i < 1440; i++){
                    var hours = Math.trunc(i/60);
                    var minutes = i % 60;
                    labelLoadHourMin.push(hours.toString()+":"+minutes.toString())
                    dataLoad.push(dataAll[requestSelect][i].totalcount);
                  }
                  
                  this.LoadChartLine(labelLoadHourMin,dataLoad,'#số lần gọi: ', 'lần')
                } else if (requestMinute === 1) {
                  //var keyData = requestSelect.toString();
                  var dataAll = request_data_hours_line;
      
                  var dataLoad = [];
                  var i = 0;
                  for (i = 0; i < 24; i++){
                    dataLoad.push(dataAll[requestSelect][i].totalcount);
                  }
                  
                  this.LoadChartLine(labelLoadHour,dataLoad,'#số lần gọi: ','lần')
                } else if (requestMinute === 3) {
                 
                  var dataAll = request_data_minutes;
      
                  var dataLoad = [];
                  var labelLoadHourMin = [];
                  var i = 0;
                  for (i = 0; i < 1440; i++){
                    var hours = Math.trunc(i/60);
                    var minutes = i % 60;
                    labelLoadHourMin.push(hours.toString()+":"+minutes.toString())
                    dataLoad.push(dataAll[requestSelect][i].biggesttime);
                  }
                  
                  this.LoadChartLine(labelLoadHourMin,dataLoad,'#thời gian xử lí cao nhất: ','nanosecond')
                }
              } 
              else {
                if (requestSelect === ''){
                  requestSelect = all_action[0]
                }
                if(requestMinute === 2) {
                  //var keyData = requestSelect.toString();
                  var dataAll = action_data_minutes;
      
                  var dataLoad = [];
                  var labelLoadHourMin = [];
                  var i = 0;
                  for (i = 0; i < 1440; i++){
                    var hours = Math.trunc(i/60);
                    var minutes = i % 60;
                    labelLoadHourMin.push(hours.toString()+":"+minutes.toString())
                    dataLoad.push(dataAll[requestSelect][i].totalcount);
                  }
                  
                  this.LoadChartLine(labelLoadHourMin,dataLoad,'#số lần gọi: ','lần')
                } else if (requestMinute === 1) {
                
                  var dataAll = action_data_hours_line;
                  console.log(dataAll)
                  console.log(requestSelect)
                  var dataLoad = [];
                  var i = 0;
                  for (i = 0; i < 24; i++){
                    dataLoad.push(dataAll[requestSelect][i].totalcount);
                  }
                  
                  this.LoadChartLine(labelLoadHour,dataLoad,'#số lần gọi: ','lần')
                } else if(requestMinute === 3) {
                
                  var dataAll = action_data_minutes;
      
                  var dataLoad = [];
                  var labelLoadHourMin = [];
                  var i = 0;
                  for (i = 0; i < 1440; i++){
                    var hours = Math.trunc(i/60);
                    var minutes = i % 60;
                    labelLoadHourMin.push(hours.toString()+":"+minutes.toString())
                    dataLoad.push(dataAll[requestSelect][i].biggesttime);
                  }
                  
                  this.LoadChartLine(labelLoadHourMin,dataLoad,'#thời gian xử lí cao nhất: ','nanosecond')
                } 
              }
          }
          
          // onload
          console.log(requestSelect);
          window.onload = function (){
             this.CheckSelectRequest(); 
             this.CheckLoadCharBar();
          };
        
    </script>
  </body>
</html>