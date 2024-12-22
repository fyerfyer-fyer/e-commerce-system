import React from 'react';
import { Form, Input, Button, Typography, message } from 'antd';
import { useNavigate, Link } from 'react-router-dom';
import { useAuth } from '../../utils/auth';
import { UserOutlined, LockOutlined } from '@ant-design/icons';

const { Title } = Typography;

function LoginPage() {
  const [form] = Form.useForm();
  const navigate = useNavigate();
  const { login } = useAuth();

  const onFinish = (values) => {
    if (values.username === 'admin' && values.password === 'password') {
      login({
        id: 1,
        username: 'admin',
        role: 'admin',
      });
      message.success('Login successful!');
      navigate('/user/info');
    } else {
      message.error('Invalid username or password.');
    }
  };

  return (
    <div className="login-container">
      <div className="login-card">
        <Title level={2} className="login-title">Welcome Back</Title>
        <Form
          form={form}
          name="login"
          onFinish={onFinish}
          layout="vertical"
          size="large"
          className="login-form"
        >
          <Form.Item
            name="username"
            rules={[{ required: true, message: 'Please enter your username!' }]}
          >
            <Input
              prefix={<UserOutlined className="site-form-item-icon" />}
              placeholder="Username"
            />
          </Form.Item>
          <Form.Item
            name="password"
            rules={[{ required: true, message: 'Please enter your password!' }]}
          >
            <Input.Password
              prefix={<LockOutlined className="site-form-item-icon" />}
              placeholder="Password"
            />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" block>
              Log in
            </Button>
          </Form.Item>
          <Link to="/register" className="register-link">
            Don't have an account? Register now
          </Link>
        </Form>
      </div>
    </div>
  );
}

export default LoginPage;