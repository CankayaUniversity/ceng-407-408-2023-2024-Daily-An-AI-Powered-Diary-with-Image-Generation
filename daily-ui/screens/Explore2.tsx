import React from 'react';
import { View, StyleSheet, Text, TouchableOpacity, Image, ScrollView, Dimensions, Modal, TextInput, Button, TouchableWithoutFeedback, Keyboard } from 'react-native';
import Header from '../components/Header';
import { DailyResponse, ReportDailyRequest, getExplore, useReportDaily } from '../libs';
import { useState, useEffect } from 'react';
import { AxiosError } from 'axios';
import "react-native-get-random-values";
import { v4 as uuidv4 } from 'uuid';
import { Ionicons } from '@expo/vector-icons';

const Explore2 = ({ navigation }) => {
  const [error, setError] = useState<AxiosError | null>(null);
  const [data, setData] = useState<any[]>([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [isVisible, setVisible] = useState(true);
  const [isLoading, setLoading] = useState(false);
  const [contentOffset, setContentOffset] = useState(0);
  const [isModalVisible, setModalVisible] = useState(false);
  const [reportText, setReportText] = useState('');
  const {mutate} = useReportDaily();

  const handleSwipe = () => {
    setCurrentPage((currentPage) => currentPage + 1);
  };

  const handleDoubleTap = () => {
    console.log(contentOffset / Dimensions.get('window').height);
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
    console.log('Report submitted:', reportText);
    setModalVisible(false);
    setReportText('');
  };

  const handleModalCancel = () => {
    setModalVisible(false);
    setReportText('');
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
        pagingEnabled
        scrollEnabled={isVisible}
        onMomentumScrollBegin={({ nativeEvent }) => {
          setVisible(true);
        }}>
        {data?.length != 0 && data?.map((el, index) => {
          return (
            <View key={uuidv4()} style={{ height: Dimensions.get('window').height, width: Dimensions.get('window').width, opacity: 1.0, backgroundColor: '#0D1326' }}>
              <View style={{ height: '100%', width: '100%' }}>
                {
                  isVisible &&
                  <Image source={{ uri: el.image }} style={styles.image}></Image>
                }
                {
                  !isVisible &&
                  <ScrollView scrollEnabled={true}>
                    <Text style={styles.text}>{el.text}</Text>
                  </ScrollView>
                }
                <TouchableOpacity style={{ position: "absolute", right: 0, top: "35%" }}>
                  <Ionicons name="heart" size={48} color="red" />
                </TouchableOpacity>
                <TouchableOpacity style={{ position: "absolute", right: 0, top: "45%" }} onPress={handleReportPress}>
                  <Ionicons name="flag" size={48} color="white" />
                </TouchableOpacity>
                <TouchableOpacity style={{ width: '90%', backgroundColor: "purple", position: "absolute", left: "5%", top: "80%", borderRadius: 8, opacity: !isVisible ? 0.3 : 1 }} onPress={() => setVisible(!isVisible)}>
                  <Text style={{ textAlign: 'center', marginBottom: 10, marginTop: 10, fontSize: 20, fontWeight: '200', color: 'white' }}>{isVisible ? "Show Image" : "Show Text"}</Text>
                </TouchableOpacity>
              </View>
            </View>
          )
        })}
      </ScrollView>
      <Modal
        visible={isModalVisible}
        transparent={true}
        animationType="slide"
        onRequestClose={() => setModalVisible(false)}
      >
        <View style={styles.modalOverlay}>
          <TouchableWithoutFeedback onPress={Keyboard.dismiss} style={{height:"100%",width:"100%"}}>
            <View style={styles.modalContent}>
              <Text style={styles.modalTitle}>Report</Text>
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
  }
});

export default Explore2;
