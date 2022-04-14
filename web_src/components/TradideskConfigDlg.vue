<template>
    <FormDlg title="桌面配置" @hide="onHide" @show="onShow" @submit="onSubmit" ref="dlg" :disabled="errors.any() || bLoading">     
         <div :class="['form-group', { 'has-error': errors.has('deskname')}]">
            <label for="input-desk-name" class="col-sm-3 control-label">桌面名称</label>
            <div class="col-sm-8">
                <input type="text" id="input-desk-name" class="form-control" name="deskname" data-vv-as="桌面名称" v-model.trim="form.deskname">
                <span class="help-block">{{errors.first('videoname')}}</span>
            </div>
        </div>
        <div :class="['form-group', { 'has-error': errors.has('deskloc')}]">
            <label for="input-desk-loc" class="col-sm-3 control-label">桌面位置</label>
            <div class="col-sm-8">
                <input type="text" id="input-desk-loc" class="form-control" name="deskloc" data-vv-as="桌面位置" v-model.trim="form.deskloc">
                <span class="help-block">{{errors.first('deskloc')}}</span>
            </div>
        </div>
        
       <div :class="['form-group', { 'has-error': errors.has('username')}]">
            <label for="input-username" class="col-sm-3 control-label">用户名</label>
            <div class="col-sm-8">
                <input type="text" id="input-username" class="form-control" name="username" data-vv-as="用户名" v-model.trim="form.username">
                <span class="help-block">{{errors.first('username')}}</span>
            </div>
        </div>
        <div :class="['form-group', { 'has-error': errors.has('deskkey')}]">
            <label for="input-deskkey" class="col-sm-3 control-label">关键字</label>
            <div class="col-sm-8">
                <input type="text" id="input-deskkey" class="form-control" name="deskkey" data-vv-as="关键字" v-model.trim="form.deskkey">
                <span class="help-block">{{errors.first('deskkey')}}</span>
            </div>
        </div>
        <div :class="['form-group', { 'has-error': errors.has('deskdesc')}]">
            <label for="input-deskdesc" class="col-sm-3 control-label">桌面描述</label>
            <div class="col-sm-8">
                <input type="text" id="input-deskdesc" class="form-control" name="deskdesc" data-vv-as="桌面描述" v-model.trim="form.deskdesc">
                <span class="help-block">{{errors.first('deskdesc')}}</span>
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
                deskname: '',
                deskloc: '',
                username: '',
                deskkey: '',
                deskdesc: '',
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
            $.get('/api/v1/tradidesk/tradideskconfig', this.form).then(data => {
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
