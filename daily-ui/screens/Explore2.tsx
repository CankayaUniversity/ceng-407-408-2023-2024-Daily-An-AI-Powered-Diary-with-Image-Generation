import React from 'react';
import { View, StyleSheet, Text, TouchableOpacity, Image, ScrollView, Dimensions, Modal, TextInput, Button, TouchableWithoutFeedback, Keyboard } from 'react-native';
import Header from '../components/Header';
import { DailyResponse, ReportDailyRequest, getExplore, useReportDaily } from '../libs';
import { useState, useEffect } from 'react';
import { AxiosError } from 'axios';
import "react-native-get-random-values";
import { v4 as uuidv4 } from 'uuid';
import { Ionicons, FontAwesome } from '@expo/vector-icons';
import { Dropdown } from 'react-native-element-dropdown';

const Explore2 = ({ navigation }) => {
  const [error, setError] = useState<AxiosError | null>(null);
  const [data, setData] = useState<any[]>([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [isVisible, setVisible] = useState(true);
  const [isLoading, setLoading] = useState(false);
  const [contentOffset, setContentOffset] = useState(0);
  const [isModalVisible, setModalVisible] = useState(false);
  const [reportText, setReportText] = useState('');
  const [selectedCategory, setSelectedCategory] = useState("");
  const { mutate } = useReportDaily();

  const reportCategories = [
    { label: "Inappropriate Content", value: "Inappropriate Content" },
    { label: "Privacy Violations", value: "Privacy Violations" },
    { label: "Spam and Scams", value: "Spam and Scams" },
    { label: "Illegal Activities", value: "Illegal Activities" },
    { label: "Self-Harm and Suicide", value: "Self-Harm and Suicide" },
    { label: "Other Violations", value: "Other Violations" }
  ];

  const getMaxEmotion = (data:any) => {
    let maxEmotionValue = -Infinity;
    let maxEmotion = "";
 
    for (const [emotion, value] of Object.entries(data.emotions)) {
       if (value as number > maxEmotionValue) {
          maxEmotionValue= value as number;
          maxEmotion = emotion;
       }
    }
    return maxEmotion;
  }

  const handleSwipe = () => {
    setCurrentPage((currentPage) => currentPage + 1);
  };

  useEffect(() => {
    const abortController = new AbortController();
    const fetchData = async () => {
      try {
        const newData = await getExplore();
        setData(data => [...data, ...newData]);
        setError(null);
        setLoading(true);
      } catch (error: any) {
        setError(error);
        console.error('Failed to fetch', error);
      }
    };
    fetchData();
    return () => abortController.abort();
  }, [currentPage]);

  useEffect(() => {
    if (error) {
      const axiosError = error as AxiosError;
      if (axiosError.response?.status === 401) {
        console.log("Unauthorized, redirecting to login");
        navigation.navigate('Login');
      }
    }
  }, [error, navigation]);

  const isCloseToBottom = ({ layoutMeasurement, contentOffset, contentSize }) => {
    const paddingToBottom = 20;
    return layoutMeasurement.height + contentOffset.y >= contentSize.height - paddingToBottom;
  };

  const handleReportPress = () => {
    setModalVisible(true);
  };

  const handleModalSubmit = () => {
    console.log('Report submitted:', selectedCategory, reportText);
    const reportDaily:ReportDailyRequest={
      dailyId: data[contentOffset / Dimensions.get('window').height].id,
      content: reportText,
      title: selectedCategory
    }
    mutate(reportDaily);
    setModalVisible(false);
    setReportText('');
    setSelectedCategory("");
  };

  const handleModalCancel = () => {
    setModalVisible(false);
    setReportText('');
    setSelectedCategory("");
  };

  return (
    <Header navigation={navigation} previous="Home" homepage={false}>
      <ScrollView
        onScroll={({ nativeEvent }) => {
          setContentOffset(nativeEvent.contentOffset.y);
          if (isCloseToBottom(nativeEvent)) {
            handleSwipe();
          }
        }}
        showsVerticalScrollIndicator={false}
        snapToInterval={Dimensions.get('screen').height}
        decelerationRate="fast"
        scrollEnabled={isVisible}
        onMomentumScrollBegin={({ nativeEvent }) => {
          setVisible(true);
        }}>
        {data?.length !== 0 && data?.map((el, index) => {
          return (
            <View key={uuidv4()} style={{ height: Dimensions.get('screen').height, width: Dimensions.get('screen').width, opacity: 1.0, backgroundColor: '#0D1326' }}>
              <View style={{ height: '100%', width: '100%' }}>
                {isVisible && <Image source={{ uri: el.image }} style={styles.image}></Image>}
                {!isVisible && (
                  <ScrollView scrollEnabled={true}>
                    <Text style={styles.text}>{el.text}</Text>
                  </ScrollView>
                )}
                <TouchableOpacity style={{position:"absolute",top:0,left:5,borderWidth:1,alignItems:"center",justifyContent:"center", aspectRatio: 2 / 1, width: '35%', opacity: 0.95, marginTop: 10, borderRadius: 10, backgroundColor: '#2D1C40' }}>
                    <Text style={styles.cardText}>{getMaxEmotion(el.emotions)}</Text>
                </TouchableOpacity>
                <TouchableOpacity style={{position:"absolute",top:0,right:5,alignItems:"center",justifyContent:"center",borderWidth:1, aspectRatio: 2 / 1, width: '35%', opacity: 0.95, marginTop: 10, borderRadius: 10, backgroundColor: '#2D1C40' }}>
                    <Text style={styles.cardText}>{"Topic".toUpperCase()}</Text>
                </TouchableOpacity>
                <TouchableOpacity style={styles.heartButton}>
                  <Ionicons name="heart" size={48} color="white" />
                </TouchableOpacity>
                <TouchableOpacity style={styles.flagButton} onPress={handleReportPress}>
                  <Ionicons name="flag" size={48} color="white" />
                </TouchableOpacity>
                <TouchableOpacity style={styles.refreshButton} onPress={() => setVisible(!isVisible)}>
                  <FontAwesome name="refresh" color="white" size={64} />
                </TouchableOpacity>
              </View>
            </View>
          );
        })}
      </ScrollView>
      <Modal
        visible={isModalVisible}
        transparent={true}
        animationType="slide"
        onRequestClose={() => setModalVisible(false)}
      >
        <View style={styles.modalOverlay}>
          <TouchableWithoutFeedback onPress={Keyboard.dismiss}>
            <View style={styles.modalContent}>
              <Text style={styles.modalTitle}>Report</Text>
              <Dropdown
                style={styles.dropdown}
                data={reportCategories}
                labelField="label"
                valueField="value"
                placeholder="Select a category"
                value={selectedCategory}
                onChange={item => {
                  setSelectedCategory(item.value);
                }}
              />
              <TextInput
                style={styles.textInput}
                placeholder="Enter your report"
                placeholderTextColor="#999"
                value={reportText}
                onChangeText={setReportText}
                enablesReturnKeyAutomatically
                multiline={true}
              />
              <View style={styles.modalButtons}>
                <Button title="Cancel" onPress={handleModalCancel} />
                <Button title="Submit" onPress={handleModalSubmit} />
              </View>
            </View>
          </TouchableWithoutFeedback>
        </View>
      </Modal>
    </Header>
  );
};

const styles = StyleSheet.create({
  text: {
    textAlign: 'left',
    paddingLeft: 10,
    paddingRight: 10,
    paddingBottom: 30,
    fontSize: 20,
    fontWeight: '200',
    color: 'white'
  },
  cardText: {
    textAlign: 'center',
    fontSize: 25,
    fontWeight: '200',
    color: 'white'
 },
  image: {
    resizeMode: 'contain',
    paddingTop: 570
  },
  modalOverlay: {
    flex: 1,
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
    justifyContent: 'center',
    alignItems: 'center'
  },
  modalContent: {
    width: '80%',
    backgroundColor: 'white',
    borderRadius: 10,
    padding: 20,
    alignItems: 'center'
  },
  modalTitle: {
    fontSize: 20,
    marginBottom: 20
  },
  dropdown: {
    width: '100%',
    marginBottom: 20,
    borderColor: '#ccc',
    borderWidth: 1,
    borderRadius: 5,
    padding: 10
  },
  textInput: {
    width: '100%',
    height: 100,
    borderColor: '#ccc',
    borderWidth: 1,
    borderRadius: 5,
    padding: 10,
    marginBottom: 20,
    textAlignVertical: 'top'
  },
  modalButtons: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    width: '100%'
  },
  heartButton: {
    width: 64,
    height: 64,
    position: 'absolute',
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: 32,
    right: Dimensions.get('screen').width / 2 + 64,
    bottom: 98
  },
  flagButton: {
    width: 64,
    height: 64,
    position: 'absolute',
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: 32,
    left: Dimensions.get('screen').width / 2 + 64,
    bottom: 98
  },
  refreshButton: {
    width: 80,
    height: 80,
    position: 'absolute',
    alignItems: 'center',
    justifyContent: 'center',
    left: Dimensions.get('screen').width / 2 - 40,
    bottom: 90,
    borderRadius: 40
  }
});

export default Explore2;