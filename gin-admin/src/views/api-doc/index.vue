<template>
    <div class="tab-container">
        <el-container>
            <el-header></el-header>
            <el-container>
              <el-main width="55%">
                    <el-tabs :tab-position="tabPosition">
                        <el-tab-pane label="Response">
                            <el-container>
                                请求返回信息
                            </el-container>
                        </el-tab-pane>
                        <el-tab-pane label="Request">
                                <el-container>
                                    请求信息
                                </el-container>
                        </el-tab-pane>
                    </el-tabs>
              </el-main>
              <el-aside width="45%">
                    <el-row :gutter="20">
                        <el-col :span="4">
                                <el-select v-model="value" placeholder="请选择">
                                        <el-option
                                              v-for="item in options"
                                              :key="item.value"
                                              :label="item.label"
                                              :value="item.value">
                                        </el-option>
                                </el-select>
                        </el-col>
                        <el-col :span="16">
                            <el-input v-model="input" placeholder="请输入内容"></el-input>
                        </el-col>
                    </el-row>
                    <el-row :gutter="20">
                        Query Parameters
                    </el-row>
                    <el-row :gutter="20">
                        <el-col :span="20">
                                <el-table
                                ref="multipleTable"
                                :data="tableData"
                                tooltip-effect="dark"
                                style="width: 100%"
                                @selection-change="handleSelectionChange">
                                <el-table-column
                                  type="selection"
                                  width="55">
                                </el-table-column>
                                <el-table-column
                                  label="key"
                                  width="120">
                                  <template slot-scope="scope">{{ scope.row.date }}</template>
                                </el-table-column>
                                <el-table-column
                                  prop="name"
                                  label="value"
                                  width="120">
                                </el-table-column>
                              </el-table>
                        </el-col>
                    </el-row>
                    <el-row :gutter="20">
                        Body
                    </el-row>
                    <el-row :gutter="20">
                        <el-col :span="20">
                            <el-input
                                    type="textarea"
                                    :rows="2"
                                    placeholder="请输入内容"
                                    v-model="textarea">
                            </el-input>
                        </el-col>
                    </el-row>
                    <el-row :gutter="20">
                            Headers
                    </el-row>
                    <el-row :gutter="20">
                            <el-col :span="20">
                                    <el-table
                                    ref="multipleTable"
                                    :data="tableData"
                                    tooltip-effect="dark"
                                    style="width: 100%"
                                    @selection-change="handleSelectionChange">
                                    <el-table-column
                                      type="selection"
                                      width="55">
                                    </el-table-column>
                                    <el-table-column
                                      label="日期"
                                      width="120">
                                      <template slot-scope="scope">{{ scope.row.date }}</template>
                                    </el-table-column>
                                    <el-table-column
                                      prop="name"
                                      label="姓名"
                                      width="120">
                                    </el-table-column>
                                    <el-table-column
                                      prop="address"
                                      label="地址"
                                      show-overflow-tooltip>
                                    </el-table-column>
                                  </el-table>
                            </el-col>
                    </el-row>
                    <el-row :gutter="20">
                        <el-col>
                            <el-button type="primary" round>执行</el-button>
                        </el-col>
                    </el-row>
              </el-aside>
            </el-container>
          </el-container>
    </div>
</template>
    
<style>
    .el-header {
      background-color: #4c5264;
      color: #333;
      text-align: center;
      /* line-height: 60px; */
      height: 10%;
      font-size: 14px;
    }
    
    .el-aside {
      background-color: #2d2f3b;
      color: white;
      text-align: center;
    }
    
    .tab-container {
        height: 100%;
    }
    
    .el-main {
      background-color: transparent;
      color: #333;
      text-align: center;
    }
    
    .el-container {
        height: 100%;
    }
    
    .el-select {
        background-color: #4c5264;
    }

    .el-col {
        margin-top: 20px;
        margin-bottom: 20px;
    }

    .el-input input  {
        background-color: #4c5264;
        border:none;
        outline:medium;
    }

    .el-button {
        background-color: inherit;
        width: 30%;
        border:2px solid #4c5264; 
    }

    .el-button:hover {
      background-color: inherit;
      border:2px solid white; 
    }

    .el-button:active {
        background-color: inherit;
        border:2px solid white; 
    }

    textarea {
        background-color: #4c5264;
    }
</style>
  