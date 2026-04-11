# ==========================================
# Stage 1: Dependencies (Install node_modules)
# ==========================================
FROM node:18-alpine AS deps
WORKDIR /app

# Menggunakan npm ci untuk instalasi yang lebih konsisten dan cepat
COPY package.json package-lock.json* ./
RUN npm ci

# ==========================================
# Stage 2: Builder (Compile Next.js)
# ==========================================
FROM node:18-alpine AS builder
WORKDIR /app

COPY --from=deps /app/node_modules ./node_modules
COPY . .

# Matikan Next.js telemetry saat build
ENV NEXT_TELEMETRY_DISABLED=1

# Build aplikasi
RUN npm run build

# ==========================================
# Stage 3: Runner (Production Environment)
# ==========================================
FROM node:18-alpine AS runner
WORKDIR /app

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1

# Hanya copy file yang dibutuhkan untuk production dari stage builder
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

# Expose port Next.js
EXPOSE 3000
ENV PORT=3000

# Jalankan server production
CMD ["npm", "start"]