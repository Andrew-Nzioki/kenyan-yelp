# KENYA-YELP
- Core User Functionality (mvp)

    1. Users can register and create profiles
    2. Businesses can be listed with full details
    3. Reviews and ratings system
    4. Search by location and categories
    5. Business owner responses
    6. Photo uploads
    7. Bookmarking favorite places
    8. Follow other users
    9. Business claiming system


- Frontend-Features:

    1. Search interface with filters
    2. Maps integration for location-based search
    3. Business profile pages
    4. Review writing interface
    5. Photo upload system
    6. User profile pages
    7. Mobile responsive design

- Backend-Services:
    1. Authentication system
    2. Search engine (Elasticsearch or algolia)
    3. Geolocation services
    4. Email notification system
    5. Content moderation system
    6. CDN for media storage
    7. Caching layer (Redis/Memcached)
    8. API layer

- Additional-Features:

    1. Reservations system
    2. Waitlist management
    3. Menu integration
    4. Q&A section
    5. Business events/deals
    6. Check-in system
    7. Review analytics
    8. Business analytics dashboard



## Database schema models

1. **Users**
- Stores user account information
- Contains: id, email, password_hash, name, role (user/owner/admin), status
- Used for: Authentication, profile management, access control
- Critical fields: email (unique), role (for permissions)

2. **Locations**
- Stores address and coordinate information
- Shared by both users and businesses
- Contains: street_address, city, state, postal_code, latitude, longitude
- Used for: Geo-searches, address management
- Separated from businesses table for better organization and reusability

3. **Businesses**
- Core business listing information
- Contains: name, description, owner_id, location_id, status, price_range
- Tracks: average_rating, review_count
- Used for: Main business listings and searches
- Links to: owner (user), location, categories

4. **Categories**
- Structured category system
- Contains: name, slug, parent_id (for hierarchy), level, status
- Used for: Business classification, search/filtering
- Enables: Hierarchical categorization (e.g., Restaurants → Italian → Pizza)

5. **Business_Categories** (Junction Table)
- Links businesses to categories (many-to-many)
- Contains: business_id, category_id
- Used for: Category-based searching and filtering

6. **Business_Hours**
- Store operating hours
- Contains: business_id, day_of_week, open_time, close_time, is_closed
- Used for: Displaying when businesses are open/closed
- Separate table for easier updates and queries

7. **Reviews**
- Store user reviews
- Contains: business_id, user_id, rating, content, status, helpful_votes
- Used for: Business ratings and user feedback
- Triggers business rating recalculation

8. **Review_Responses**
- Store business owner responses to reviews
- Contains: review_id, user_id, content
- Used for: Owner engagement with customers
- One response per review (unique constraint)

9. **Media**
- Handles all uploaded media
- Contains: entity_id, entity_type, media_type (image/video), url
- Used for: Photos/videos for businesses, reviews, users
- Generic design allows attachment to different entities

10. **Business_Claims**
- Manages business verification process
- Contains: business_id, user_id, status, verification_documents
- Used for: Business ownership verification
- Tracks claim status and approval process

11. **Bookmarks**
- Stores user's saved businesses
- Contains: user_id, business_id
- Used for: User's saved/favorite places
- Simple junction table with timestamps

12. **User_Follows**
- Manages user relationships
- Contains: follower_id, following_id
- Used for: Social features
- Junction table for user-to-user relationships

13. **Review_Votes**
- Tracks helpful/unhelpful votes on reviews
- Contains: user_id, review_id, vote_type
- Used for: Review quality assessment
- Prevents duplicate voting (unique constraint)

Key Design Points:
1. All tables have created_at timestamps
2. Most tables have status fields for soft deletion
3. UUIDs used for IDs to prevent enumeration
4. Foreign keys properly constrained
5. Appropriate indexes on frequently queried fields

This structure provides:
- Clean separation of concerns
- Efficient querying capabilities
- Scalable design for future features
- Data integrity through constraints
- Proper relationship management