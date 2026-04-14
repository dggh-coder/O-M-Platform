<template>
  <el-container style="min-height: 100vh">
    <el-aside width="240px" style="border-right: 1px solid #e5e7eb; padding: 16px">
      <h3>O&M Platform</h3>
      <el-menu router>
        <el-menu-item index="/dashboard">Dashboard</el-menu-item>
        <el-menu-item index="/hosts">Hosts</el-menu-item>
        <el-menu-item index="/packs">Packs</el-menu-item>
        <el-menu-item index="/projects">Projects</el-menu-item>
        <el-menu-item index="/deployments">Deployments</el-menu-item>
        <el-menu-item index="/runtime">Runtime</el-menu-item>
        <el-menu-item index="/metrics">Metrics</el-menu-item>
        <el-menu-item index="/settings">Settings</el-menu-item>
      </el-menu>
    </el-aside>
    <el-main><RouterView /></el-main>
  </el-container>
</template>
