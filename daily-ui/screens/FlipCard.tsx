import React, { useState, useRef } from 'react';
import { View, ScrollView, Text, Image, Pressable, Animated, StyleSheet,Dimensions } from 'react-native';
import uuidv4 from 'uuid/v4';

interface FlipCardProps {
  dailyUrl: string;
  dailyContent: string;
};

const FlipCard = ({ dailyUrl, dailyContent }: FlipCardProps) => {
  const [isFlipped, setIsFlipped] = useState(false); const animatedValue = useRef(new Animated.Value(0)).current;
  const frontInterpolate = animatedValue.interpolate({
    inputRange: [0, 180],
    outputRange: ['0deg', '180deg'],
  });
  const backInterpolate = animatedValue.interpolate({
    inputRange: [0, 180],
    outputRange: ['180deg', '360deg'],
  });

  const frontAnimatedStyle = {
    transform: [{ rotateY: frontInterpolate }],
  };

  const backAnimatedStyle = {
    transform: [{ rotateY: backInterpolate }],
  };

  const toggleFlip = () => {
    const toValue = isFlipped ? 0 : 180;

    Animated.spring(animatedValue, {
      toValue: toValue,
      friction: 8,
      tension: 10,
      useNativeDriver: false,
    }).start();

    setIsFlipped(!isFlipped);
  };

  return (
    <Pressable onPress={toggleFlip}>
      <View>
        <Animated.View style={[styles.flipCard, frontAnimatedStyle, { opacity: isFlipped ? 0 : 1 }]}>
          <Animated.Image source={{ uri: dailyUrl }} style={styles.image} />
        </Animated.View>
        <Animated.View style={[styles.flipCard, styles.flipCardBack, backAnimatedStyle, { opacity: isFlipped ? 1 : 0 }]}>
          <ScrollView contentContainerStyle={styles.flipCardBackInside}>
            <Text style={styles.textStyle}>{dailyContent}</Text>
          </ScrollView>
        </Animated.View>
      </View>
    </Pressable>
  );
};

const styles = StyleSheet.create({
  textStyle: {
    fontSize: 20,
    color: 'white',
  },
  flipCard: {
    backgroundColor: 'black',
    justifyContent: 'center',
    backfaceVisibility: 'hidden',
  },
  flipCardBack: {
    position: 'absolute',
    width: '100%',
    height: '100%',
  },
  flipCardBackInside: {
    fontSize: 15,
    opacity: 0.85,
    backgroundColor: '#0D1326',
    height: '100%',
    width: '100%',
  },
  image: {
    resizeMode: 'contain',
    width: '100%',
    height: '100%',
  },
});

export default FlipCard;
