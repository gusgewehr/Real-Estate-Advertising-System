import 'dart:io';

import 'package:dio/dio.dart';
import 'package:flutter/foundation.dart';
import 'package:frontend/core/api_client.dart';
import 'package:frontend/data/models/real_estate.dart';
import 'package:file_picker/file_picker.dart';

import '../../domain/entity/pagination.dart';
import '../../domain/entity/real_estate.dart';
import '../models/pagination.dart';

abstract class RealEstateDataSource{
  Future<void> createRealEstate(RealEstateEntity realEstate);
  Future<PaginationEntity<RealEstateEntity>> getRealEstate(int page, int pageSize);
  Future<String> uploadImage(PlatformFile image);
}

class RealEstateDataSourceImpl implements RealEstateDataSource{
  final ApiClient apiClient;

  RealEstateDataSourceImpl({required this.apiClient});

  @override
  Future<void> createRealEstate(RealEstateEntity realEstate) async {
    try{

      var json = RealEstateModel.fromEntity(realEstate).toJSON();

      final response = await apiClient.dio.post('/real-estate', data: json);


    }on DioException catch (e) {
      throw Exception('Connection Failed: ${e.message}');}
    catch(e){
      throw Exception('Error fetching data: $e');
    }
  }

  @override
  Future<PaginationEntity<RealEstateEntity>> getRealEstate(int page, int pageSize ) async {
    try{

      final Map<String, dynamic> params = {
        'page': page,
        'pageSize': pageSize,
        };

      final response = await apiClient.dio.get('/real-estate', queryParameters: params);

      if (response.statusCode != 200) {
        throw Exception('Failed to fetch real estate: ${response.statusCode}');
      }


      return PaginationModel<RealEstateEntity>.fromJson(response.data, RealEstateModel.fromJSON);



    } on DioException catch (e) {
      throw Exception('Connection Failed: ${e.message}');}
    catch(e){
      throw Exception('Error fetching data: $e');
    }


  }

  @override
  Future<String> uploadImage(PlatformFile pickedFile) async {
    try {
      MultipartFile multipartFile;

      if (kIsWeb) {
        multipartFile = MultipartFile.fromBytes(
          pickedFile.bytes!,
          filename: pickedFile.name,
        );
      } else {
        multipartFile = await MultipartFile.fromFile(
          pickedFile.path!,
          filename: pickedFile.name,
        );
      }

      FormData formData = FormData.fromMap({
        "image": multipartFile
      });

      Response response = await apiClient.dio.post(
        "/real-estate/image",
        data: formData,
      );

      return response.data;
    } on DioException catch (e) {
      throw Exception('Connection Failed: ${e.message}');
    } catch(e){
      throw Exception('Error fetching data: $e');
    }


    return "";
  }


}