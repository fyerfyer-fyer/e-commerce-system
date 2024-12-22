import React from 'react';
import { Layout } from 'antd';

const { Footer } = Layout;

function AppFooter() {
  return (
    <Footer style={{ textAlign: 'center', backgroundColor: '#001529', color: 'white' }}>
      © 2023 E-Commerce Platform. All rights reserved.
    </Footer>
  );
}

export default AppFooter;