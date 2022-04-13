<template>
    <div class="container-fluid no-padding">

        <div class="box box-success">
            <div class="box-header">
                <h4 class="text-success text-center">传统桌面列表</h4>           
                <form class="form-inline">
                  <div class="form-group">
                      <button type="button" class="btn btn-sm btn-success" @click.prevent="$refs['TradideskConfigDlg'].show()"><i class="fa fa-plus"></i>添加传统桌面</button>
                  </div>                  
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
                <el-table :data="pushers" stripe class="view-list" :default-sort="{prop: 'startAt', order: 'descending'}" @sort-change="sortChange">
                    <el-table-column prop="Id" label="ID" min-width="160"></el-table-column>
                    <el-table-column prop="Deskname" label="桌面名称" min-width="80"></el-table-column>
                    <el-table-column prop="Deskloc" label="桌面位置" min-width="80"></el-table-column>
                    <el-table-column prop="Username" label="用户名称" min-width="80"></el-table-column>
                    <el-table-column prop="Deskdesc" label="桌面描述" min-width="240"></el-table-column>
                    <el-table-column prop="Deskkey" label="桌面关键字" min-width="240"></el-table-column>               
                    <el-table-column label="操作" min-width="120" fixed="right">
                        <template slot-scope="scope">
                            <div class="btn-group">
                                <a role="button" class="btn btn-xs btn-danger" @click.prevent="stop(scope.row)">
                                  <i class="fa fa-stop"></i> 删除
                                </a>
                            </div>
                        </template>
                    </el-table-column>
                </el-table>          
            </div>
            <div class="box-footer clearfix" v-if="total > 0">
                <el-pagination layout="prev,pager,next" class="pull-right" :total="total" :page-size.sync="pageSize" :current-page.sync="currentPage"></el-pagination>
            </div>
        </div>     
        <TradideskConfigDlg ref="TradideskConfigDlg" @submit="getTradideskconfs"></TradideskConfigDlg>          
    </div>
</template>

<script>
import TradideskConfigDlg from "components/TradideskConfigDlg.vue"
import prettyBytes from "pretty-bytes";

import _ from "lodash";
export default {
  components: {
    TradideskConfigDlg
  },
  props: {},
  data() {
    return {
      q: "",
      sort: "startAt",
      order: "descending",
      pushers: [],
      total: 0,
      timer: 0,
      pageSize: 10,
      currentPage: 1
    };
  },
  beforeDestroy() {
    if (this.timer) {
      clearInterval(this.timer);
      this.timer = 0;
    }
  },
  mounted() {
    // this.$refs["q"].focus();
    this.getTradideskconfs();
    this.timer = setInterval(() => {
      this.getTradideskconfs();
    }, 3000);
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
    getTradideskconfs() {
      $.get("/api/v1/tradidesk/tradidesklist", {
        q: this.q,
        start: (this.currentPage -1) * this.pageSize,
        limit: this.pageSize,
        sort: this.sort,
        order: this.order
      }).then(data => {
        this.total = data.total;
        this.pushers = data.rows;
      });
    },
    doSearch(page = 1) {
      var query = {};
      if (this.q) query["q"] = this.q;
      this.$router.replace({
        path: `/tradideskconf/${page}`,
        query: query
      });
    },
    doDelaySearch: _.debounce(function() {
      this.doSearch();
    }, 500),    
    sortChange(data) {
      this.sort = data.prop;
      this.order = data.order;
      this.getVideoconfs();
    },
    formatBytes(row, col, val) {
      if (val == undefined) return "-";
      return prettyBytes(val);
    },
    stop(row) {
      this.$confirm(`确认删除 ${row.Deskname} ?`, "提示").then(() => {
        $.get("/api/v1/tradidesk/remove", {
          id: row.Id
        }).then(data => {
          this.getTradideskconfs();
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
      this.pushers = [];
      this.getTradideskconfs();
    });
  }
};
</script>

