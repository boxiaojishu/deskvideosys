<template>
<div class="box box-primary videos">
    <div class="box-header">
        <h4 class="text-primary text-center">桌面行为记录</h4>
    </div>
    <div class="box-body">
        <div class="form-inline">
            <div class="form-group">
                <div class="btn-group">
                    <button type="button" class="btn btn-primary btn-sm" @click.prevent="$router.go(-1)">
                        <i class="fa fa-chevron-left"></i> 返回
                    </button>
                    <button type="button" class="btn btn-danger btn-sm" @click.prevent="removeDaily" v-if="videos.length > 0">
                        <i class="fa fa-remove"></i> 删除当天
                    </button>
                </div>
            </div>
            <div class="form-group pull-right">
                <div class="input-group input-group-sm">
                    <WindowRecordDatePicker class="form-control" :day="day" @update:day="updateDay" ref="datePicker" :user="user" :virname="virname"></WindowRecordDatePicker>
                    <div class="input-group-btn">
                        <button type="button" class="btn btn-sm btn-default" @click.prevent="showWindowRecordDatePicker">
                            <i class="fa fa-calendar"></i>
                        </button>
                        <router-link :to="`/windowrecord/timeview/${this.user}/${this.virname}/${this.day}`" replace class="btn btn-default btn-sm">
                            <i class="fa fa-hand-o-right"></i> 时间轴视图
                        </router-link>
                    </div>
                </div>
            </div>
        </div>
        <br>
        <div class="clearfix"></div>
        <el-table :row-class-name="tableRowClassName" ref="recordTable" :data="pageData" empty-text="暂无数据, 请另选日期" :default-sort="{prop: 'startAt', order: 'descending'}" @sort-change="sortChange">
            <el-table-column min-width="120" label="用户名称" prop="user" show-overflow-tooltip></el-table-column>
            <el-table-column min-width="120" label="虚拟机名称" prop="virname" show-overflow-tooltip></el-table-column>
            <el-table-column min-width="200" label="应用名称" prop="appname" sortable="custom"></el-table-column>
            <el-table-column min-width="200" label="窗口名称" prop="winname" sortable="custom"></el-table-column>
            <el-table-column min-width="120" label="开始时间" prop="startAt" sortable="custom"></el-table-column>
            
            <el-table-column min-width="120" label="运行时长" prop="duration" sortable="custom"></el-table-column>
        </el-table>
    </div>
    <div class="box-footer clearfix" v-if="total > 0">
        <el-pagination layout="prev,pager,next, jumper" class="pull-right" :total="total" :page-size.sync="pageSize" :current-page.sync="currentPage"></el-pagination>
    </div>

</div>
</template>

<script>
import {
    mapState
} from "vuex";
import moment from "moment";
import WindowRecordDatePicker from "components/WindowRecordDatePicker.vue";
import axios from "axios";
import _ from 'lodash'

export default {
    props: {
        user: {type:String,default:'zhongguo'},
        virname: {type:String,default:'hunan'},
        day: String
    },
    data() {
        return {
            q: "",
            name: "",
            currentPage: 1,
            pageSize: 10,
            total: 0,
            sort: "startAt",
            order: "ascending",
            videos: [],
            bgcolor: "",
            pathname: location.pathname == "/" ? "" : location.pathname.substring(0, location.pathname.length - 1)
        };
    },
    components: {
        WindowRecordDatePicker
    },
    methods: {
        tableRowClassName({
            row,
            rowIndex
        }) {
            if (row.important) {
                return "warning-row";
            }
        },
        isMobile() {
            return videojs.browser.IS_IOS || videojs.browser.IS_ANDROID;
        },
        updateDay(day) {
            this.$router.replace(`/windowrecordlist/${this.user}/${this.virname}/${day}`);
        },
        showWindowRecordDatePicker() {
            $(this.$refs.datePicker.$el).focus();
        },
        updateVideos() {
            $.get("/api/v1/windowrecord/querydaily", {
                user: this.user,
                virname: this.virname,
                period: this.day,
                q: this.q,
                start: (this.currentPage -1) * this.pageSize,
                limit: this.pageSize,
                sort: this.sort,
                order: this.order
            }).then(data => {
                this.videos = data.rows;
                this.total = data.total;
            });
        },
        updateVideosDebounce: _.debounce(function () {
            this.updateVideos()
        }, 500),
        play(row) {
            this.$refs["deskRecordVideoDlg"].play(
                row.hls,
                `${this.name}(${row.startAt})`,
                row.Snap || "",
            );
        },
        download(row) {
            window.open(`/api/v1/cloudrecord/download/${this.user}/${this.virname}/${row._startAt}`);
        },
        removeDaily() {
            this.$confirm(`确认删除 ${this.name} ${moment(this.day).format("YYYY-MM-DD")} 当天所有录像?`,"提示")
                .then(() => {
                    $.get("/api/v1/cloudrecord/removedaily", {
                        user: this.user,
                        virname: this.virname,
                        period: this.day
                    }).always(() => {
                        this.updateVideos();
                        this.$refs["datePicker"].update();
                    });
                })
                .catch(() => {});
        },
        remove(row) {
            this.$confirm(`确认删除 ${row.name} ?`, "提示")
                .then(() => {
                    $.get("/api/v1/windowrecord/remove", {
                        user: this.user,
                        virname: this.virname,
                        day: this.day,
                        period: row._startAt
                    }).always(() => {
                        this.updateVideos();
                    });
                })
                .catch(() => {});
        },
        sortChange(data) {
            this.sort = data.prop;
            this.order = data.order;
            this.updateVideos();
        },
        turnImportant(row) {
            $.get("/api/v1/windowrecord/setimportant", {
                user: row.user,
                virname: row.virname,
                day: this.day,
                period: row._startAt,
                important: row.important
            }).then(() => {
                var msg = row.important ? "标记成功,将不会被自动清理" : "取消标记成功";
                this.$message({
                    type: "success",
                    message: msg
                });
            })
        }
    },
    mounted() {
        if (!this.user || !this.virname) {
            this.$router.replace("/cloudesks");
            return;
        }
        if (!this.day) {
            this.$router.replace(`/videodlg/${this.user}/${this.virname}/${moment().format("YYYYMMDD")}`);
            return;
        }
         this.updateVideos();
    },
    computed: {
        ...mapState(["serverInfo", "userInfo"]),
        total() {
            return this.videos.length;
        },
        pageData() {
            let start = (this.currentPage - 1) * this.pageSize;
            let end = start + this.pageSize;
            return this.videos.slice(start, end).map(row => {
                if (!row.user) {
                    row.user = this.user,
                    row.virname = this.virname,
                    row.name = this.name,
                    row.location = location;
                    row._startAt = row.startAt;
                    row.startAt = moment(row.startAt, "YYYYMMDDHHmmss").format("YYYY-MM-DD HH:mm:ss");
                    row.duration = moment.utc((row.duration || 0) * 1000).format("HH:mm:ss");
                }
                return row;
            });
        }
    },
    beforeRouteUpdate(to, from, next) {
         if (!to.params.user || !to.params.virname) {
            this.$router.replace(`/cloudesks`);
            return;
        }
        
        if (!to.params.day) {
            this.$router.replace(`/videodlg/${to.params.user}/${to.params.virname}/${moment().format("YYYYMMDD")}`);
            return;
        }
        next();
        this.$nextTick(() => {
            this.updateVideos();
        });
    }
};
</script>

<style>
.el-table .warning-row {
    background: oldlace;
    color: red;
}

.el-table .success-row {
    background: #f0f9eb;
}
</style>


