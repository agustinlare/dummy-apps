package com.clave.dummy.controller

import com.amazonaws.services.s3.model.*
import com.fasterxml.jackson.databind.node.ObjectNode
import org.slf4j.LoggerFactory
import org.springframework.beans.factory.annotation.Value
import org.springframework.http.HttpStatus.*
import org.springframework.http.MediaType
import org.springframework.http.MediaType.*
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController
import reactor.core.publisher.Mono

@RestController
@RequestMapping("/properties-tester")
class PropertiesTestController(

    @Value("\${dummy.property.value}")
    private val propertyValue: String,

) {

    @GetMapping(
        path = ["/print-dummy-property"],
        produces = [APPLICATION_JSON_VALUE]
    )
    fun printVariable() : Mono<Map<String, Any>> {
        logger.info("Printing dummy.property.value")
        return Mono.just(mapOf("dummy_property_value" to propertyValue))
    }


    companion object {
        @Suppress("JAVA_CLASS_ON_COMPANION")
        @JvmStatic
        private val logger = LoggerFactory.getLogger(javaClass.enclosingClass)
    }

}