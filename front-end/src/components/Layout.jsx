import React from "react";
import { Outlet, Link } from "react-router-dom";
import { Layout, Menu } from "antd";

const { Header, Content, Footer } = Layout;

const AppLayout = () => {
  return (
    <Layout className="layout">
      <Header>
        <div className="logo" style={{ color: "white", fontSize: "20px" }}>
          E-Commerce
        </div>
        <Menu theme="dark" mode="horizontal" defaultSelectedKeys={["1"]}>
          <Menu.Item key="1">
            <Link to="/">首页</Link>
          </Menu.Item>
          <Menu.Item key="2">
            <Link to="/cart">购物车</Link>
          </Menu.Item>
          <Menu.Item key="3">
            <Link to="/order">订单</Link>
          </Menu.Item>
          <Menu.Item key="4">
            <Link to="/user">个人中心</Link>
          </Menu.Item>
          <Menu.Item key="5">
            <Link to="/seckill">秒杀活动</Link>
          </Menu.Item>
        </Menu>
      </Header>
      <Content style={{ padding: "20px 50px" }}>
        <Outlet />
      </Content>
      <Footer style={{ textAlign: "center" }}>
        E-Commerce ©2024 Created by ChatGPT
      </Footer>
    </Layout>
  );
};

export default AppLayout;