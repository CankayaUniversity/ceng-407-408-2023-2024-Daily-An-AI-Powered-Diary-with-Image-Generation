import React from 'react';
import { useState, useEffect } from 'react';
import { View, StyleSheet, Text, Image } from 'react-native';
import Header from '../components/Header';
import { Calendar } from 'react-native-calendars';
import { Colors } from '../libs/colors.tsx';
import uuidv4 from 'uuid/v4';
import { getStatistics } from '../libs/services/dailyService';
import { StatisticsResponse } from '../libs/types';

const calendarSelectedColor = '#AF5C5C66';

const Statistics = ({ navigation }: { navigation: any }) => {
   const [timeframe, setTimeframe] = useState("month");
   const [statistics, setStatistics] = useState<StatisticsResponse | null>(null);
   const [isLoading, setIsLoading] = useState<boolean>(true);
   const [error, setError] = useState<string | null>(null);
   const [markedDatesState, setMarkedDatesState] = useState<any>(null);

   useEffect(() => {
      const fetchData = async () => {
         try {
            const data = await getStatistics();
            setStatistics(data);
         } catch (err) {
            setError("Failed to fetch statistics");
            console.error(err);
         } finally {
            setIsLoading(false);
         }
      };

      fetchData();
   }, []);


   const markedDatesFunc = statistics?.date.reduce((acc: any, date) => {
      acc[date] = { selected: true };
      return acc;
   }, {});

   useEffect(() => {
      console.log(markedDatesFunc);
      setMarkedDatesState(markedDatesFunc);
   }, [statistics]);


   return (
      <Header navigation={navigation} previous="Home" homepage={false}>
         <View style={styles.container}>
            <Calendar
               key={uuidv4()}

               markedDates={
                  markedDatesState
               }

               style={{
                  backgroundColor: Colors.main_container,
                  width: '100%',
                  paddingLeft: 0,
                  paddingRight: 0
               }}

               theme={{
                  textDayFontSize: 20,
                  textMonthFontSize: 20,
                  selectedDayTextColor: 'white',
                  calendarBackground: 'transparent',
                  monthTextColor: Colors.background,
                  dayTextColor: 'white',
                  'stylesheet.calendar.header': {
                     fontSize: 20,
                     width: '100%',
                     alignItems: 'center',
                     marginTop: 5,
                     flexDirection: 'row',
                     justifyContent: 'center',
                     dayHeader: {
                        marginBottom: 7,
                        width: 36,
                        textAlign: 'center',
                        fontSize: 16,
                        color: "white"
                     },
                  },
                  'stylesheet.calendar.main': {
                     dayContainer: {
                        flex: 1,
                        alignItems: 'center'
                     },
                     week: {
                        marginVertical: 3,
                        marginHorizontal: 0,
                        flexDirection: 'row',
                        justifyContent: 'center'
                     },
                  },
                  'stylesheet.day.basic': {
                     selected: {
                        color: calendarSelectedColor,
                        backgroundColor: calendarSelectedColor,
                        width: 40,
                        height: 36,
                        borderRadius: 10,
                     },
                  }
               }
               }
               hideArrows={true}
            />

            <View style={styles.outerRow}>
               <View style={styles.innerItem}>
                  <View>
                     <Text style={styles.innerText}>
                        Wows
                     </Text>
                     <View style={{ width: "100%", height: 34 }}></View>
                     <View style={{ flexDirection: "row", alignItems: "center" }} >
                        <Image source={require('../assets/Heart.png')} style={{ marginRight: 10 }} />
                        <Text style={styles.innerNumber}>
                           {statistics?.likes || "0"}
                        </Text>
                     </View>
                  </View>
               </View>
               <View style={styles.innerItem}>
                  <View>
                     <Text style={styles.innerText}>
                        Dailies written
                     </Text>
                     <View style={{ width: "100%", height: 34 }}></View>
                     <View style={{ flexDirection: "row", alignItems: "center" }} >
                        <Image source={require('../assets/increase.png')} style={{ marginRight: 10, marginTop: 5 }} />
                        <Text style={styles.innerNumber}>
                           {statistics?.dailiesWritten || 0}
                        </Text>
                     </View>
                  </View>
               </View>
            </View>
            <View style={styles.outerRow}>
               <View style={styles.innerItem}>
                  <View>
                     <Text style={styles.innerText}>
                        Views
                     </Text>
                     <View style={{ width: "100%", height: 34 }}></View>
                     <View style={{ flexDirection: "row", alignItems: "center" }} >
                        <Image source={require('../assets/increase.png')} style={{ marginRight: 10, marginTop: 5 }} />
                        <Text style={styles.innerNumber}>
                           {statistics?.views || 0}
                        </Text>
                     </View>
                  </View>
               </View>
               <View style={styles.innerItem}>
                  <View>
                     <Text style={styles.innerText}>
                        Streak
                     </Text>
                     <View style={{ width: "100%", height: 34 }}></View>
                     <View style={{ flexDirection: "row", alignItems: "center" }} >
                        <Image source={require('../assets/streak.png')} style={{ marginRight: 10, marginTop: 5 }} />
                        <Text style={styles.innerNumber}>
                           {statistics?.streak || 0}
                        </Text>
                     </View>
                  </View>
               </View>
            </View>
            <View style={styles.outerRow}>
               <View style={styles.innerItem}>
                  <View>
                     <Text style={styles.innerText}>
                        Mood
                     </Text>
                     <View style={{ width: "100%", height: 20 }}></View>
                     <View style={{ flexDirection: "row", width: "70%", height: "56%", alignItems: "flex-end" }} >
                        <Image source={require('../assets/happy.png')} style={{ marginRight: 10, }} />
                        <Text style={[styles.innerNumber, { fontSize: 20, fontWeight: "bold" }]}>
                           {statistics?.mood || "Uncalculated"}
                        </Text>
                     </View>
                  </View>
               </View>
               <View style={styles.innerItem}>
                  <View>
                     <Text style={styles.innerText}>
                        Topic
                     </Text>
                     <View style={{ width: "100%", height: 20 }}></View>
                     <View style={{ flexDirection: "row", width: "70%", height: "56%", alignItems: "flex-end" }} >
                        <Text style={[styles.innerNumber, { fontSize: 20, fontWeight: "bold" }]}>
                           {statistics?.topic || "Uncalculated"}
                        </Text>
                     </View>
                  </View>
               </View>
            </View>
         </View>
      </Header >
   );
}

const styles = StyleSheet.create({
   container: {
      flex: 1, alignItems: 'center',
      padding: 10,
      margin: 10,
      borderRadius: 20,
      backgroundColor: Colors.main_container,
      opacity: 0.90,
   },
   outerRow: {
      height: '15%',
      marginTop: 20,
      width: '92%',
      flexDirection: 'row',
      justifyContent: 'space-between',
   },
   innerItem: {
      padding: 10,
      borderRadius: 10,
      width: '45%',
      marginHorizontal: 10,
      backgroundColor: "rgba(0,0,0,0.6)",
      flexDirection: "column",
      justifyContent: "space-between",
      height: '100%',
   },
   innerText: {
      fontFamily: "Helvetica",
      fontSize: 20,
      color: "rgba(255, 255, 255, 1)",
   },
   innerNumber: {
      marginTop: 4,
      fontSize: 40,
      color: "white",
   }

})
export default Statistics;
