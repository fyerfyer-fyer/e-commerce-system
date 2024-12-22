import React from 'react';
import { Menu, Layout } from 'antd';

const { Sider } = Layout;

function AppSidebar() {
  return (
    <Sider width={200} style={{ background: '#fff' }}>
      <Menu mode="inline" defaultSelectedKeys={['1']} style={{ height: '100%' }}>
        <Menu.Item key="1">Option 1</Menu.Item>
        <Menu.Item key="2">Option 2</Menu.Item>
        <Menu.Item key="3">Option 3</Menu.Item>
      </Menu>
    </Sider>
  );
}

export default AppSidebar;