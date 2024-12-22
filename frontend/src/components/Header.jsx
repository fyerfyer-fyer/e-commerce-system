import React from 'react';
import { Link } from 'react-router-dom';
import { Menu, Layout } from 'antd';

const { Header } = Layout;

function AppHeader() {
  return (
    <Header style={{ backgroundColor: '#001529' }}>
      <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['/']}>
        <Menu.Item key="/">
          <Link to="/">Home</Link>
        </Menu.Item>
        <Menu.Item key="/login">
          <Link to="/login">Login</Link>
        </Menu.Item>
        <Menu.Item key="/register">
          <Link to="/register">Register</Link>
        </Menu.Item>
        <Menu.Item key="/user/info">
          <Link to="/user/info">User Info</Link>
        </Menu.Item>
        <Menu.Item key="/cart/list">
          <Link to="/cart/list">Cart</Link>
        </Menu.Item>
        <Menu.Item key="/order/submit">
          <Link to="/order/submit">Submit Order</Link>
        </Menu.Item>
        <Menu.Item key="/payment/make">
          <Link to="/payment/make">Make Payment</Link>
        </Menu.Item>
        <Menu.Item key="/comment/list/1">
          <Link to="/comment/list/1">Comments</Link>
        </Menu.Item>
        <Menu.Item key="/seckill/events">
          <Link to="/seckill/events">Seckill Events</Link>
        </Menu.Item>
        <Menu.Item key="/admin/product/add">
          <Link to="/admin/product/add">Admin Product Add</Link>
        </Menu.Item>
      </Menu>
    </Header>
  );
}

export default AppHeader;