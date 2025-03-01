FROM openjdk:17-slim

COPY ./build/libs/data-generator-all.jar /app/data-generator-all.jar

# Set the working directory
WORKDIR /app

# Run the JAR file
CMD ["java", "-jar", "data-generator-all.jar"]
