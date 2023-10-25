package main;

import java.io.FileInputStream;
import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import org.apache.poi.ss.usermodel.CellType;
import org.apache.poi.ss.util.CellAddress;
import org.apache.poi.xssf.usermodel.*;

public class Main {

	public static void main(String[] args) throws IOException {
		LinkedHashMap<String, Object> jsonMap = new LinkedHashMap<>();
		String courseName = "courseName";

		String excelFilePath = ".\\Rooms_copied.xlsx";
		FileInputStream inputStream = new FileInputStream(excelFilePath);
		XSSFWorkbook workbook = new XSSFWorkbook(inputStream);
		XSSFSheet sheet = workbook.getSheetAt(0);
		CellAddress firstRoomCell = new CellAddress("Y4");
		List<String> daysList = List.of("sunday", "monday", "tuesday", "wednesday", "thursday");
		int roomCol = firstRoomCell.getColumn();
		int lastRow = sheet.getLastRowNum();

		// loop over all the rooms
		for (int currRow = 3; currRow < lastRow; currRow++) {
			XSSFRow roomRow = sheet.getRow(currRow);
			if (roomRow.getLastCellNum() <= 1) {
				continue;
			}
			XSSFCell roomCell = roomRow.getCell(roomCol);
			if (roomCell.getCellType() != CellType.STRING || roomCell.getStringCellValue().equals("ONLINE")) {
				continue;
			}

			LinkedHashMap<String, Object> innerMap = new LinkedHashMap<>();
			innerMap.put("name", roomCell.getStringCellValue());
			// initializing each day as null the, semicolons are used to format it into json
			// map
			daysList.forEach(day -> innerMap.put(day, ";null;"));
			XSSFRow daysRow = sheet.getRow(currRow + 1);

			// check which days are busy in the particular room
			for (int j = 9; j < daysRow.getLastCellNum(); j++) {
				ArrayList<LinkedHashMap<String, Object>> dayCourses = new ArrayList<>();
				XSSFCell daysCell = daysRow.getCell(j);
				if (daysCell.getCellType() != CellType.STRING) {
					continue;
				}
				String dayCellString = daysCell.getStringCellValue().toLowerCase();

				for (int coursesRowNum = daysRow.getRowNum() + 2;; coursesRowNum++) {
					LinkedHashMap<String, Object> courseDetailsMap = new LinkedHashMap<>();
					XSSFRow coursesRow = sheet.getRow(coursesRowNum);
					if (coursesRow == null || coursesRow.getLastCellNum() <= 1) {
						break;
					}
					XSSFCell courseCell = coursesRow.getCell(j);
					if (courseCell.getCellType() != CellType.STRING
							|| coursesRow.getCell(0).getCellType() == CellType.BLANK) {
						continue;
					}
					String courseTimeStart = coursesRow.getCell(0).getStringCellValue();
					int indexOfStartColon = courseTimeStart.indexOf(':');
					String startHourString = courseTimeStart.substring(0, indexOfStartColon);
					String startMinString = courseTimeStart.substring(indexOfStartColon + 1, indexOfStartColon + 3);
					int startHour = convertTOInt(startHourString);
					int startMin = convertTOInt(startMinString);
					if (courseTimeStart.endsWith("PM") && startHour != 12) {
						startHour += 12;
					}
					LinkedHashMap<String, String> startTimesMap = new LinkedHashMap<>(
							Map.of("hour", ";" + startHour + ";"));
					startTimesMap.put("minute", ";" + startMin + ";");

					String courseTimeEnd = coursesRow.getCell(5).getStringCellValue();

					int indexOfEndColon = courseTimeEnd.indexOf(':');
					String endHourString = courseTimeEnd.substring(0, indexOfEndColon);
					String endMinString = courseTimeEnd.substring(indexOfEndColon + 1, indexOfEndColon + 3);
					int endHour = convertTOInt(endHourString);
					int endMin = convertTOInt(endMinString);
					if (courseTimeEnd.endsWith("PM") && endHour != 12) {
						endHour += 12;
					}
					LinkedHashMap<String, String> endTimesMap = new LinkedHashMap<>(
							Map.of("hour", ";" + endHour + ";"));
					endTimesMap.put("minute", ";" + endMin + ";");

					courseDetailsMap.put("timeStart", startTimesMap);
					courseDetailsMap.put("timeEnd", endTimesMap);
					courseDetailsMap.put(courseName, courseCell.getStringCellValue());
					if (dayCourses.size() > 0) {
						LinkedHashMap<String, Object> freeTimeMap = new LinkedHashMap<>();
						LinkedHashMap<String, Object> prevEntry = dayCourses.get(dayCourses.size() - 1);
						LinkedHashMap<String, String> prevTimeEnd = (LinkedHashMap<String, String>) prevEntry
								.get("timeEnd");
						freeTimeMap.put("timeStart", prevTimeEnd);
						freeTimeMap.put("timeEnd", startTimesMap);
						freeTimeMap.put("courseName", "Free");
						dayCourses.add(freeTimeMap);
					}
					dayCourses.add(courseDetailsMap);

				}
				innerMap.put(dayCellString, dayCourses);
			}
			jsonMap.put(roomCell.getStringCellValue(), innerMap);

		}
		workbook.close();

		// converting java map to json map
		String rawMapString = jsonMap.toString().replace("=", ":").replace(Character.toString(160), " ");
		StringBuilder jsonMapString = new StringBuilder();
		char[] rawMapCharArray = rawMapString.toCharArray();
		boolean isQuoteOpen = false;
		boolean isPaused = false;
		for (char character : rawMapCharArray) {
			if (!Character.isLetterOrDigit(character)) {
				if (character == ';') {
					isPaused = !isPaused;
					continue;
				}
				if (isQuoteOpen && !Character.isWhitespace(character) && character != '-' && character != '_') {
					isQuoteOpen = !isQuoteOpen;
					jsonMapString.append('"');
				}
				jsonMapString.append(character);
				continue;
			}
			if (!isQuoteOpen && !isPaused) {
				jsonMapString.append('"');
				isQuoteOpen = !isQuoteOpen;
			}
			jsonMapString.append(character);

		}
		saveStrToFile(jsonMapString.toString());

	}

	public static Integer convertTOInt(String str) {
		try {
			int res = Integer.parseInt(str);
			return res;

		} catch (Exception e) {
			System.err.println("given string is not a number");
			return null;
		}
	}

	public static void saveStrToFile(String input) {
		String fileName = "output.txt";
		try {
			FileWriter outputFile = new FileWriter(fileName);
			outputFile.write(input);
			outputFile.close();
		} catch (IOException e) {
			System.err.println("There was an error accessing the file " + fileName);
		}
	}

}
