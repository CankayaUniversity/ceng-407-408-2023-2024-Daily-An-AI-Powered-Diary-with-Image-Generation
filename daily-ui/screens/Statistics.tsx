import React from 'react';
import { View, StyleSheet, Text } from 'react-native';
import Header from '../components/Header';
import { Calendar } from 'react-native-calendars';
import { Colors } from '../libs/colors.tsx';
import uuidv4 from 'uuid/v4';

{/*  #6082B */ }

const Statistics = ({ navigation }: { navigation: any }) => {
   return (
      <Header navigation={navigation} previous="Home" homepage={false}>
         <View style={styles.container}>
            <Calendar
               key={uuidv4()}

               markedDates={{
                  '2024-05-22': { selected: true, selectedColor: 'gray' }
               }}

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
                     week: {
                        fontSize: 20,
                        width: '100%',
                        alignItems: 'center',
                        marginTop: 5,
                        flexDirection: 'row',
                        justifyContent: 'center'
                     },
                     dayHeader: {
                        marginBottom: 7,
                        width: 40,
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
                        borderRadius: 4,
                     },
                  }
               }
               }
               hideArrows={true}
            />
         </View>
      </Header>
   );
}

const styles = StyleSheet.create({
   container: {
      flex: 1, alignItems: 'center',
      padding: 10,
      margin: 10,
      borderRadius: 20,
      backgroundColor: Colors.main_container,
      opacity: 0.85,
   }
})
export default Statistics;
