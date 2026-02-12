import type { User, AuthResponse, LoginCredentials, RegisterData } from '../types';
import { storage } from '../utils/storage';

// Simulate network delay
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

// Mock auth service using localStorage
export const authService = {
  async login(credentials: LoginCredentials): Promise<AuthResponse> {
    await delay(800);

    const users = storage.getUsers();
    const user = users.find(
      (u: any) => u.email === credentials.email && u.password === credentials.password
    );

    if (!user) {
      throw new Error('Invalid email or password');
    }

    const token = `mock_token_${user.id}_${Date.now()}`;
    storage.setToken(token);

    return {
      user: {
        id: user.id,
        name: user.name,
        email: user.email,
      },
      token,
    };
  },

  async register(data: RegisterData): Promise<AuthResponse> {
    await delay(800);

    const users = storage.getUsers();
    
    // Check if user already exists
    if (users.find((u: any) => u.email === data.email)) {
      throw new Error('Email already registered');
    }

    const newUser = {
      id: `user_${Date.now()}`,
      name: data.name,
      email: data.email,
      password: data.password, // In real app, this would be hashed
      createdAt: new Date().toISOString(),
    };

    users.push(newUser);
    storage.setUsers(users);

    const token = `mock_token_${newUser.id}_${Date.now()}`;
    storage.setToken(token);

    return {
      user: {
        id: newUser.id,
        name: newUser.name,
        email: newUser.email,
      },
      token,
    };
  },

  async forgotPassword(email: string): Promise<{ message: string }> {
    await delay(1000);

    // Always return success message (don't reveal if email exists)
    return {
      message: 'If an account exists with this email, you will receive a password reset link.',
    };
  },

  async resetPassword(token: string, newPassword: string): Promise<{ message: string }> {
    await delay(800);

    // In a real app, we would verify the token and update the password
    // For mock purposes, we'll just simulate success
    return {
      message: 'Your password has been reset successfully.',
    };
  },

  async getCurrentUser(): Promise<User | null> {
    await delay(300);

    const token = storage.getToken();
    if (!token) {
      return null;
    }

    // Extract user ID from token (mock implementation)
    const userIdMatch = token.match(/mock_token_([^_]+)/);
    if (!userIdMatch) {
      return null;
    }

    const userId = userIdMatch[1];
    const users = storage.getUsers();
    const user = users.find((u: any) => u.id === userId);

    if (!user) {
      return null;
    }

    return {
      id: user.id,
      name: user.name,
      email: user.email,
    };
  },

  async logout(): Promise<void> {
    await delay(300);
    storage.removeToken();
  },
};
