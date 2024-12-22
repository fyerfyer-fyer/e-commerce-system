import React from 'react';
import { Form, Input, Button, Typography, message } from 'antd';
import { Link } from 'react-router-dom';
import { useAuth } from '../../utils/auth';
import { UserOutlined, MailOutlined, LockOutlined } from '@ant-design/icons';

const { Title, Paragraph } = Typography;

function RegisterPage() {
  const [form] = Form.useForm();
  const { login } = useAuth();

  const onFinish = (values) => {
    if (values.password === values.confirmPassword) {
      const newUser = {
        id: Math.floor(Math.random() * 1000),
        username: values.username,
        email: values.email,
        role: 'user',
      };

      login(newUser);
      message.success('Registration successful! You are now logged in.');
      form.resetFields();
      window.location.href = '/user/info';
    } else {
      message.error('Passwords do not match.');
    }
  };

  return (
    <div className="register-container">
      <div className="register-card">
        <Title level={2} className="register-title">Create Account</Title>
        <Paragraph className="register-subtitle">
          Join us to get started with your shopping journey
        </Paragraph>
        <Form
          form={form}
          name="register"
          onFinish={onFinish}
          layout="vertical"
          className="register-form"
          size="large"
        >
          <Form.Item
            name="username"
            rules={[{ required: true, message: 'Please enter your username!' }]}
          >
            <Input
              prefix={<UserOutlined />}
              placeholder="Username"
            />
          </Form.Item>
          <Form.Item
            name="email"
            rules={[
              { required: true, message: 'Please enter your email!' },
              { type: 'email', message: 'Please enter a valid email address!' }
            ]}
          >
            <Input
              prefix={<MailOutlined />}
              placeholder="Email"
            />
          </Form.Item>
          <Form.Item
            name="password"
            rules={[
              { required: true, message: 'Please enter your password!' },
              { min: 6, message: 'Password must be at least 6 characters!' }
            ]}
          >
            <Input.Password
              prefix={<LockOutlined />}
              placeholder="Password"
            />
          </Form.Item>
          <Form.Item
            name="confirmPassword"
            dependencies={['password']}
            rules={[
              { required: true, message: 'Please confirm your password!' },
              ({ getFieldValue }) => ({
                validator(_, value) {
                  if (!value || getFieldValue('password') === value) {
                    return Promise.resolve();
                  }
                  return Promise.reject(new Error('The two passwords do not match!'));
                },
              }),
            ]}
          >
            <Input.Password
              prefix={<LockOutlined />}
              placeholder="Confirm Password"
            />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" block>
              Register
            </Button>
          </Form.Item>
          <Link to="/login" className="login-link">
            Already have an account? Login now
          </Link>
        </Form>
      </div>
    </div>
  );
}

export default RegisterPage;