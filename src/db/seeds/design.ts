import { db } from "@/db/index"
import { designs, type TNewDesign } from "@/db/schema"

export async function seedDesign(input: { count: number }) {
  const count = input.count ?? 100

  try {
    const allDesigns: TNewDesign[] = []

    // Add some predefined designs with rich content
    allDesigns.push(...getPredefinedDesigns())

    // Add random designs
    for (let i = 0; i < count - 3; i++) {
      allDesigns.push(generateRandomDesign())
    }

    await db.delete(designs)

    console.log("📝 Inserting designs", allDesigns.length)

    await db.insert(designs).values(allDesigns).onConflictDoNothing()
  } catch (err) {
    console.error(err)
  }
}

export function generateRandomDesign(input?: Partial<TNewDesign>): TNewDesign {
  const designNumber = Math.floor(Math.random() * 1000)
  const sampleBodies = [
    `# Design Overview ${designNumber}

This is a sample design document with **markdown formatting**.

## Features
- Feature 1: Basic functionality
- Feature 2: Advanced options
- Feature 3: User-friendly interface

## Code Example
\`\`\`javascript
function hello() {
    console.log("Hello, World!");
}
\`\`\`

> This is a sample design for demonstration purposes.`,

    `# Project Specification

## Introduction
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.

### Technical Details
1. First requirement
2. Second requirement
3. Third requirement

\`inline code example\`

**Bold text** and *italic text* for emphasis.`,

    null, // Some designs without body content
  ]

  return {
    title: `Design ${designNumber}`,
    description: `This is a description for design ${designNumber}`,
    body: sampleBodies[Math.floor(Math.random() * sampleBodies.length)],
    ...input,
  }
}

function getPredefinedDesigns(): TNewDesign[] {
  return [
    {
      title: "User Authentication System",
      description: "Complete authentication system with JWT tokens and role-based access control",
      body: `# User Authentication System

## Overview
This document outlines the design for a comprehensive user authentication system that provides secure login, registration, and role-based access control.

## Features

### Core Authentication
- **User Registration**: Email-based registration with verification
- **User Login**: Secure login with JWT tokens
- **Password Reset**: Email-based password reset functionality
- **Session Management**: Automatic token refresh and logout

### Security Features
- Password hashing using bcrypt
- JWT token authentication
- Rate limiting for login attempts
- CSRF protection
- Input validation and sanitization

## Technical Implementation

### Database Schema
\`\`\`sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
\`\`\`

### API Endpoints

#### POST /api/auth/register
Register a new user account.

**Request Body:**
\`\`\`json
{
    "email": "user@example.com",
    "password": "securePassword123",
    "firstName": "John",
    "lastName": "Doe"
}
\`\`\`

#### POST /api/auth/login
Authenticate user and return JWT token.

**Request Body:**
\`\`\`json
{
    "email": "user@example.com",
    "password": "securePassword123"
}
\`\`\`

**Response:**
\`\`\`json
{
    "token": "jwt_token_here",
    "user": {
        "id": "user_id",
        "email": "user@example.com",
        "role": "user"
    }
}
\`\`\`

## Security Considerations

> **Important**: All passwords must be hashed using bcrypt with a minimum salt rounds of 12.

### Password Requirements
- Minimum 8 characters
- At least one uppercase letter
- At least one lowercase letter
- At least one number
- At least one special character

### Rate Limiting
- Login attempts: 5 attempts per 15 minutes
- Registration: 3 attempts per hour per IP
- Password reset: 1 attempt per 5 minutes

## Testing Strategy

1. **Unit Tests**: Test individual authentication functions
2. **Integration Tests**: Test complete authentication flows
3. **Security Tests**: Test for common vulnerabilities
4. **Load Tests**: Ensure system can handle concurrent users

---

*This document is part of the system architecture documentation.*`,
    },
    {
      title: "E-commerce Product Catalog",
      description: "Scalable product catalog system with search, filtering, and inventory management",
      body: `# E-commerce Product Catalog

## System Architecture

### Overview
The product catalog system is designed to handle millions of products with real-time search capabilities and inventory tracking.

## Database Design

### Product Entity
| Field | Type | Description |
|-------|------|-------------|
| id | UUID | Primary key |
| name | VARCHAR(500) | Product name |
| description | TEXT | Product description |
| price | DECIMAL(10,2) | Product price |
| inventory_count | INTEGER | Available inventory |
| category_id | UUID | Foreign key to categories |
| created_at | TIMESTAMP | Creation timestamp |

### Categories Hierarchy
\`\`\`
Electronics
├── Computers
│   ├── Laptops
│   ├── Desktop
│   └── Accessories
├── Mobile Devices
│   ├── Smartphones
│   └── Tablets
└── Audio
    ├── Headphones
    └── Speakers
\`\`\`

## Search Implementation

### Elasticsearch Integration
\`\`\`javascript
// Product search index mapping
const productMapping = {
  properties: {
    name: {
      type: 'text',
      analyzer: 'standard',
      fields: {
        keyword: { type: 'keyword' }
      }
    },
    description: { type: 'text' },
    price: { type: 'float' },
    category: { type: 'keyword' },
    tags: { type: 'keyword' }
  }
}
\`\`\`

### Search API
\`\`\`typescript
interface SearchParams {
  query?: string;
  category?: string;
  minPrice?: number;
  maxPrice?: number;
  page?: number;
  limit?: number;
  sortBy?: 'price' | 'name' | 'created_at';
  sortOrder?: 'asc' | 'desc';
}
\`\`\`

## Performance Optimizations

### Caching Strategy
- **Redis**: Cache popular search results for 5 minutes
- **CDN**: Cache product images and static assets
- **Database**: Implement proper indexing on search fields

### Indexing
\`\`\`sql
-- Performance indexes
CREATE INDEX idx_products_category ON products(category_id);
CREATE INDEX idx_products_price ON products(price);
CREATE INDEX idx_products_name_gin ON products USING GIN(to_tsvector('english', name));
\`\`\`

## API Endpoints

### GET /api/products/search
Search products with filters and pagination.

**Query Parameters:**
- \`q\`: Search query
- \`category\`: Category filter
- \`min_price\`: Minimum price filter
- \`max_price\`: Maximum price filter
- \`page\`: Page number (default: 1)
- \`limit\`: Items per page (default: 20)

**Example Response:**
\`\`\`json
{
  "products": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "name": "MacBook Pro 16-inch",
      "description": "Powerful laptop for professionals",
      "price": 2399.99,
      "inventory_count": 15,
      "category": "Electronics > Computers > Laptops",
      "images": ["url1", "url2"]
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 25,
    "total_items": 500,
    "items_per_page": 20
  }
}
\`\`\`

## Inventory Management

### Real-time Updates
- WebSocket connections for inventory updates
- Optimistic locking for concurrent order processing
- Automatic reorder notifications when inventory is low

### Stock Levels
- **In Stock**: > 10 items
- **Low Stock**: 1-10 items
- **Out of Stock**: 0 items
- **Pre-order**: Negative values allowed

> **Note**: Implement proper error handling for race conditions in inventory updates.

---

*Last updated: December 2024*`,
    },
    {
      title: "Mobile App UI/UX Guidelines",
      description: "Comprehensive design system and user experience guidelines for mobile applications",
      body: `# Mobile App UI/UX Guidelines

## Design Principles

### 1. Simplicity First
Keep interfaces clean and focused on the primary user task.

### 2. Thumb-Friendly Navigation
Ensure all interactive elements are within comfortable thumb reach.

### 3. Consistent Visual Language
Use consistent colors, typography, and spacing throughout the app.

## Color Palette

### Primary Colors
- **Primary Blue**: \`#007AFF\` - For primary actions and focus states
- **Secondary Gray**: \`#8E8E93\` - For secondary text and inactive states
- **Success Green**: \`#34C759\` - For success messages and positive actions
- **Warning Orange**: \`#FF9500\` - For warnings and attention-grabbing elements
- **Error Red**: \`#FF3B30\` - For errors and destructive actions

### Usage Examples
\`\`\`css
.primary-button {
  background-color: #007AFF;
  color: white;
  border-radius: 8px;
  padding: 12px 24px;
}

.secondary-text {
  color: #8E8E93;
  font-size: 14px;
}
\`\`\`

## Typography Scale

| Style | Font Size | Line Height | Weight |
|-------|-----------|-------------|--------|
| H1 | 32px | 40px | Bold |
| H2 | 24px | 32px | Bold |
| H3 | 20px | 28px | Semibold |
| Body | 16px | 24px | Regular |
| Caption | 14px | 20px | Regular |
| Small | 12px | 16px | Regular |

## Layout Guidelines

### Grid System
- **Base unit**: 8px
- **Margins**: 16px (2 units)
- **Gutters**: 8px (1 unit)
- **Card padding**: 16px (2 units)

### Spacing Scale
\`\`\`
4px   - xs (0.5 units)
8px   - sm (1 unit)
16px  - md (2 units)
24px  - lg (3 units)
32px  - xl (4 units)
48px  - 2xl (6 units)
\`\`\`

## Component Library

### Buttons

#### Primary Button
- Background: Primary blue
- Text: White
- Border radius: 8px
- Minimum height: 44px
- Minimum width: 88px

#### Secondary Button
- Background: Transparent
- Text: Primary blue
- Border: 1px solid primary blue
- Border radius: 8px

#### Text Button
- Background: Transparent
- Text: Primary blue
- No border
- Minimum height: 44px

### Form Elements

#### Text Input
\`\`\`css
.text-input {
  border: 1px solid #E5E5E7;
  border-radius: 8px;
  padding: 12px 16px;
  font-size: 16px;
  min-height: 44px;
}

.text-input:focus {
  border-color: #007AFF;
  box-shadow: 0 0 0 3px rgba(0, 122, 255, 0.1);
}
\`\`\`

## Accessibility Standards

### Touch Targets
- **Minimum size**: 44x44px
- **Recommended**: 48x48px
- **Spacing**: At least 8px between interactive elements

### Color Contrast
- **Normal text**: 4.5:1 minimum ratio
- **Large text**: 3:1 minimum ratio
- **UI elements**: 3:1 minimum ratio

### Screen Reader Support
- Provide meaningful \`aria-label\` attributes
- Use semantic HTML elements
- Ensure logical tab order

## Animation Guidelines

### Timing Functions
- **Ease out**: For elements entering the screen
- **Ease in**: For elements leaving the screen
- **Ease in-out**: For elements changing position

### Duration
- **Micro-interactions**: 150-200ms
- **Page transitions**: 300-500ms
- **Loading animations**: 1000ms+

### Examples
\`\`\`css
.fade-in {
  animation: fadeIn 200ms ease-out;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(8px); }
  to { opacity: 1; transform: translateY(0); }
}
\`\`\`

## Platform-Specific Considerations

### iOS
- Use SF Symbols for icons
- Follow iOS Human Interface Guidelines
- Support Dynamic Type
- Use system colors when appropriate

### Android
- Follow Material Design principles
- Support different screen densities
- Use Material Icons
- Implement proper elevation and shadows

## Testing Checklist

- [ ] All touch targets are at least 44px
- [ ] Color contrast meets WCAG standards
- [ ] App works in both light and dark modes
- [ ] Text scales properly with system font size
- [ ] Navigation is consistent across screens
- [ ] Loading states are implemented
- [ ] Error states are handled gracefully

---

*These guidelines should be reviewed and updated quarterly to ensure they remain current with platform updates and best practices.*`,
    },
  ]
}
