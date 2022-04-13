<template>
    <div class="container-fluid no-padding">

        <div class="box box-success">
            <div class="box-header">
                <h4 class="text-success text-center">云桌面列表</h4>
                <form class="form-inline">
                       
                  <div class="form-group pull-right">
                      <div class="input-group">
                          <input type="text" class="form-control" placeholder="搜索" v-model.trim="q" @keydown.enter.prevent ref="q">
                          <div class="input-group-btn">
                              <button type="button" class="btn btn-default" @click.prevent="doSearch" >
                                  <i class="fa fa-search"></i>
                              </button>  
                          </div>                            
                      </div>
                  </div>                              
                </form>            
            </div>
            <div class="box-body">
                <el-table :data="cloudesks" stripe class="view-list" :default-sort="{prop: 'startAt', order: 'descending'}" @sort-change="sortChange">
                    <el-table-column prop="Virid" label="ID" min-width="200"></el-table-column>
                    <el-table-column prop="Username" label="用户名" min-width="100"></el-table-column>                            
                    <el-table-column prop="Virname" label="虚拟机名" min-width="120"  sortable="custom"></el-table-column>
                    <el-table-column label="操作" min-width="120" fixed="right">
                        <template slot-scope="scope">
                            <div class="btn-group">


                            </div>
                        </template>
                    </el-table-column>
                </el-table>          
            </div>
            <div class="box-footer clearfix" v-if="total > 0">
                <el-pagination layout="prev,pager,next" class="pull-right" :total="total" :page-size.sync="pageSize" :current-page.sync="currentPage"></el-pagination>
            </div>
        </div>
    </div>
</template>

<script>
import prettyBytes from "pretty-bytes"
import moment from 'moment';
import Vue from 'vue'
import {
    mapState
} from 'vuex'

import _ from "lodash";
export default {
  components: {
  },
  props: {},
  data() {
    return {
      q: "",
      sort: "startAt",
      order: "descending",
      cloudesks: [],
      total: 0,
      timer: 0,
      pageSize: 10,
      currentPage: 1,
      nowtime:"",
      newtime:"",
      newmonth:"",
      newyear:"",
      savedays:1
    };
  },
  beforeDestroy() {
    if (this.timer) {
      clearInterval(this.timer);
      this.timer = 0;
    }
  },
  created(){
			this.getTime();
                        this.newtime = moment().format('YYYYMMDD');
                        this.newmonth = moment().format('YYYYMM');
                        this.newyear = moment().format('YYYY');
		},
  mounted() {
    this.getCloudesks();
    this.$refs["q"].focus();
    this.timer = setInterval(() => {
      this.getCloudesks();
    }, 3000);
    this.newtime = moment().format('YYYYMMDD');
    this.newmonth = moment().format('YYYYMM');
    this.newyear = moment().format('YYYY');
  },
  watch: {
    q: function(newVal, oldVal) {
      this.doDelaySearch();
    },
    currentPage: function(newVal, oldVal) {
      this.doSearch(newVal);
    }
  },  
  methods: {
    getCloudesks() {
      $.get("/api/v1/deskcontrols", {
        q: this.q,
        start: (this.currentPage -1) * this.pageSize,
        limit: this.pageSize,
        sort: this.sort,
        order: this.order
      }).then(data => {
        this.total = data.total;
        this.cloudesks = data.rows;
      });
    },
    getTime(){
				setInterval(()=>{
					//new Date() new一个data对象，当前日期和时间
					//toLocaleString() 方法可根据本地时间把 Date 对象转换为字符串，并返回结果。
					this.nowtime = new Date().toLocaleString();
				},1000)
			},
    doSearch(page = 1) {
      var query = {};
      if (this.q) query["q"] = this.q;
      this.$router.replace({
        path: `/deskcontrols/${page}`,
        query: query
      });
    },
    doDelaySearch: _.debounce(function() {
      this.doSearch();
    }, 500),    
    sortChange(data) {
      this.sort = data.prop;
      this.order = data.order;
      this.getCloudesks();
    },
    formatBytes(row, col, val) {
      if (val == undefined) return "-";
      return prettyBytes(val);
    },
    stop(row) {
      this.$confirm(`确认停止 ${row.path} ?`, "提示").then(() => {
        $.get("/api/v1/stream/stop", {
          id: row.id
        }).then(data => {
          this.getCloudesks();
        })
      }).catch(() => {});
    }
  },
  beforeRouteEnter(to, from, next) {
    next(vm => {
      vm.q = to.query.q || "";
      vm.currentPage = parseInt(to.params.page) || 1;
    });
  },
  beforeRouteUpdate(to, from, next) {
    next();
    this.$nextTick(() => {
      this.q = to.query.q || "";
      this.currentPage = parseInt(to.params.page) || 1;
      this.cloudesks = [];
      this.getCloudesks();
    });
  }
};
</script>

