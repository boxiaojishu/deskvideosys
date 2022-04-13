<template>
    <FormDlg title="视频设备配置" @hide="onHide" @show="onShow" @submit="onSubmit" ref="dlg" :disabled="errors.any() || bLoading">     
         <div :class="['form-group', { 'has-error': errors.has('videoName')}]">
            <label for="input-video-name" class="col-sm-3 control-label">视频名称</label>
            <div class="col-sm-8">
                <input type="text" id="input-video-name" class="form-control" name="videoName" data-vv-as="视频名称" v-model.trim="form.videoName">
                <span class="help-block">{{errors.first('videoName')}}</span>
            </div>
        </div>
        <div :class="['form-group', { 'has-error': errors.has('videoLoc')}]">
            <label for="input-video-loc" class="col-sm-3 control-label">视频位置</label>
            <div class="col-sm-8">
                <input type="text" id="input-video-loc" class="form-control" name="videoLoc" data-vv-as="视频位置" v-model.trim="form.videoLoc">
                <span class="help-block">{{errors.first('videoLoc')}}</span>
            </div>
        </div>
        
        <div :class="['form-group',{'has-error': errors.has('accesstype')}]">
            <label for="input-sms-id" class="col-sm-3 control-label">接入类型
            </label>
            <div class="col-sm-8">
                <select class="form-control" id="input-accesstype" name="accesstype" v-model.trim="form.accesstype" data-vv-as="接入类型" v-validate="">
                    <option value="rtsp">RTSP</option>
                    <option value="onvif">ONVIF</option>
                </select>
            </div>
        </div>
       <div :class="['form-group', { 'has-error': errors.has('username')}]">
            <label for="input-username" class="col-sm-3 control-label">用户名</label>
            <div class="col-sm-8">
                <input type="text" id="input-username" class="form-control" name="username" data-vv-as="用户名" v-model.trim="form.username">
                <span class="help-block">{{errors.first('username')}}</span>
            </div>
        </div>
       <div :class="['form-group', { 'has-error': errors.has('userpwd')}]">
            <label for="input-userpwd" class="col-sm-3 control-label">密码</label>
            <div class="col-sm-8">
                <input type="text" id="input-userpwd" class="form-control" name="userpwd" data-vv-as="密码" v-model.trim="form.userpwd">
                <span class="help-block">{{errors.first('userpwd')}}</span>
            </div>
        </div>


        <div :class="['form-group', { 'has-error': errors.has('url')}]">
            <label for="input-url" class="col-sm-3 control-label"><span class="text-red">*</span> 地址</label>
            <div class="col-sm-8">
                <input type="text" id="input-url" class="form-control" name="url" data-vv-as="地址" v-validate="'required'" v-model.trim="form.url">
                <span class="help-block">{{errors.first('url')}}</span>
            </div>
        </div>                   
        <div class="form-group">
            <label for="input-custom-path" class="col-sm-3 control-label">录像信息</label>
            <div class="col-sm-8">
               <el-checkbox  size="mini" v-model.trim="form.isRecord" name="IsRecord">是否录像</el-checkbox>
            </div> 
        </div>
        <div :class="['form-group', { 'has-error': errors.has('savedays')}]">
           <label for="input-savedays" class="col-sm-3 control-label">录像时长</label>
           <div class="col-sm-8">
              <input type="text" id="input-savedays" class="form-control" name="savedays" data-vv-as="录像时长" v-validate="'numeric'" v-model.trim="form.savedays" placeholder="0天">
              <span class="help-block">{{errors.first('savedays')}}</span>
            </div>
         </div>
 

          
    </FormDlg>
</template>

<script>
import Vue from 'vue'
import FormDlg from 'components/FormDlg.vue'
import $ from 'jquery'

export default {
    data() {
        return {
            bLoading: false,
            form: this.defForm()
        }
    },
    components: {
        FormDlg
    },
    methods: {
        defForm() {
            return {
                videoName: '',
                videoLoc: '',
                url: '',
                accesstype: '',
                username: '',
                userpwd: '',
                customPath: '',
                transType: 'TCP',
                idleTimeout: '',
                heartbeatInterval: ''
            }
        },
        onHide() {
            this.errors.clear();
            this.form = this.defForm();
        },
        onShow() {
            document.querySelector(`[name=url]`).focus();
        },
        async onSubmit() {
            var ok = await this.$validator.validateAll();
            if(!ok) {
                var e = this.errors.items[0];
                this.$message({
                    type: 'error',
                    message: e.msg
                });
                document.querySelector(`[name=${e.field}]`).focus();
                return;
            }
            this.bLoading = true;
            $.get('/api/v1/stream/start', this.form).then(data => {
                this.$refs['dlg'].hide();
                this.$emit('submit');
            }).always(() => {
                this.bLoading = false;
            })
        },
        show(data) {
            this.errors.clear();
            if(data) {
                Object.assign(this.form, data);
            }
            this.$refs['dlg'].show();
        }
    }
}
</script>
